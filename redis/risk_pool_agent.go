package redis

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// Risk Pool Redis Keys
// Key 格式: risk_prevent_pool_agent:{type}:{AGENT}:{CURRENCY}:{LEVEL}:{GAME_NAME}
func GetRiskPoolAgentAvailableKey(agentAccount, gameName, currency string, level int) string {
	gameName = strings.ToUpper(gameName)
	if level <= 0 {
		level = 1 // 默認等級 1
	}
	return fmt.Sprintf("risk_prevent_pool_agent:amount:%s:%s:%d:%s", agentAccount, currency, level, gameName)
}

func GetRiskPoolAgentResvAmtKey(agentAccount, gameName, currency string, level int) string {
	gameName = strings.ToUpper(gameName)
	if level <= 0 {
		level = 1
	}
	return fmt.Sprintf("risk_prevent_pool_agent:resv:amt:%s:%s:%d:%s", agentAccount, currency, level, gameName)
}

func GetRiskPoolAgentResvZKey(agentAccount, gameName, currency string, level int) string {
	gameName = strings.ToUpper(gameName)
	if level <= 0 {
		level = 1
	}
	return fmt.Sprintf("risk_prevent_pool_agent:resv:exp:%s:%s:%d:%s", agentAccount, currency, level, gameName)
}

func GetRiskPoolAgentSettledSetKey(agentAccount, gameName, currency string, level int) string {
	gameName = strings.ToUpper(gameName)
	if level <= 0 {
		level = 1
	}
	return fmt.Sprintf("risk_prevent_pool_agent:settled:%s:%s:%d:%s", agentAccount, currency, level, gameName)
}

// RiskPoolAgentReserveResult represents the result of reserve operation
type RiskPoolAgentReserveResult struct {
	Status         int // 0=insufficient, 1=already_reserved, 2=reserved_ok
	ReservedAmount float64
	AvailableNow   float64 // for status=0, shows current available
	GCCount        int64   // how many expired reservations were cleaned
}

// RiskPoolAgentSettleResult represents the result of settle operation
type RiskPoolAgentSettleResult struct {
	Status           int     // 0=error, 2=success, 3=already_settled, 4=settled_without_reservation
	DeltaReturn      float64 // how much was returned to pool
	HouseGainAdded   float64 // how much house_gain was added
	OriginalReserved float64 // original reserved amount
}

// RiskPoolAgentReserve executes the reserve Lua script
// Returns: result, error
func RiskPoolAgentReserve(ctx context.Context, agentAccount, gameName, currency, betID string, amount float64, ttl time.Duration, gcLimit int64, level int) (*RiskPoolAgentReserveResult, error) {
	availableKey := GetRiskPoolAgentAvailableKey(agentAccount, gameName, currency, level)
	resvAmtKey := GetRiskPoolAgentResvAmtKey(agentAccount, gameName, currency, level)
	resvZKey := GetRiskPoolAgentResvZKey(agentAccount, gameName, currency, level)

	fmt.Println(fmt.Sprintf("agentAccount:%s, gameName:%s, currency:%s, betID:%s ", agentAccount, gameName, currency, betID))
	fmt.Println(fmt.Sprintf("availableKey:%s, resvAmtKey:%s, resvZKey:%s, betID:%s ", availableKey, resvAmtKey, resvZKey))

	nowMs := time.Now().UnixMilli()
	expireAtMs := nowMs + ttl.Milliseconds()

	// Execute Lua script
	res, err := redisConn.EvalSha(
		ctx,
		riskPoolReserveSHA,
		[]string{availableKey, resvAmtKey, resvZKey},
		betID, amount, nowMs, expireAtMs, gcLimit,
	).Result()
	if err != nil {
		return nil, fmt.Errorf("redis reserve script error: %w", err)
	}

	// Parse result
	arr, ok := res.([]interface{})
	if !ok || len(arr) < 3 {
		return nil, fmt.Errorf("invalid reserve script result: %v", res)
	}

	status, ok1 := arr[0].(int64)
	value, ok2 := parseFloat(arr[1])
	gcCount, ok3 := arr[2].(int64)

	if !ok1 || !ok2 || !ok3 {
		return nil, fmt.Errorf("invalid reserve script result types: %v", arr)
	}

	result := &RiskPoolAgentReserveResult{
		Status:  int(status),
		GCCount: gcCount,
	}

	switch status {
	case 0: // insufficient
		result.AvailableNow = value
	case 1: // already reserved
		result.ReservedAmount = value
	case 2: // reserved ok
		result.ReservedAmount = value
	}

	return result, nil
}

// RiskPoolAgentSettle executes the settle Lua script
func RiskPoolAgentSettle(ctx context.Context, agentAccount, gameName, currency, betID, result string, actualPayout, houseGain float64, level int) (*RiskPoolAgentSettleResult, error) {
	availableKey := GetRiskPoolAgentAvailableKey(agentAccount, gameName, currency, level)
	resvAmtKey := GetRiskPoolAgentResvAmtKey(agentAccount, gameName, currency, level)
	resvZKey := GetRiskPoolAgentResvZKey(agentAccount, gameName, currency, level)
	settledSetKey := GetRiskPoolAgentSettledSetKey(agentAccount, gameName, currency, level)

	// Execute Lua script
	res, err := redisConn.EvalSha(
		ctx,
		riskPoolSettleSHA,
		[]string{availableKey, resvAmtKey, resvZKey, settledSetKey},
		betID, result, actualPayout, houseGain,
	).Result()
	if err != nil {
		return nil, fmt.Errorf("redis settle script error: %w", err)
	}

	// Parse result
	arr, ok := res.([]interface{})
	if !ok || len(arr) < 4 {
		return nil, fmt.Errorf("invalid settle script result: %v", res)
	}

	status, ok1 := arr[0].(int64)
	deltaReturn, ok2 := parseFloat(arr[1])
	gainAdded, ok3 := parseFloat(arr[2])
	originalReserved, ok4 := parseFloat(arr[3])

	if !ok1 || !ok2 || !ok3 || !ok4 {
		return nil, fmt.Errorf("invalid settle script result types: %v", arr)
	}

	return &RiskPoolAgentSettleResult{
		Status:           int(status),
		DeltaReturn:      deltaReturn,
		HouseGainAdded:   gainAdded,
		OriginalReserved: originalReserved,
	}, nil
}

// GetRiskPoolAgentAvailable gets the current available pool amount
func GetRiskPoolAgentAvailable(ctx context.Context, agentAccount, gameName, currency string, level int) (float64, error) {
	key := GetRiskPoolAgentAvailableKey(agentAccount, gameName, currency, level)
	val, err := redisConn.Get(ctx, key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return 0, nil
		}
		return 0, err
	}
	return parseFloatString(val)
}

// SetRiskPoolAgentAvailable sets the available pool amount (for initialization/admin)
func SetRiskPoolAgentAvailable(ctx context.Context, agentAccount, gameName, currency string, amount float64, level int) error {
	key := GetRiskPoolAgentAvailableKey(agentAccount, gameName, currency, level)
	return redisConn.Set(ctx, key, amount, 0).Err()
}

// CleanupExpiredReservesAgent 清理過期預扣
func CleanupExpiredReservesAgent(ctx context.Context, agentAccount, gameName, currency string, limit int, level int) (cleaned int, returned float64, err error) {
	availableKey := GetRiskPoolAgentAvailableKey(agentAccount, gameName, currency, level)
	resvAmtKey := GetRiskPoolAgentResvAmtKey(agentAccount, gameName, currency, level)
	resvZKey := GetRiskPoolAgentResvZKey(agentAccount, gameName, currency, level)

	nowMs := time.Now().UnixMilli()

	result, err := redisConn.EvalSha(ctx, GetRiskPoolCleanupSHA(),
		[]string{availableKey, resvAmtKey, resvZKey},
		nowMs, limit,
	).Result()
	if err != nil {
		return 0, 0, err
	}

	if results, ok := result.([]interface{}); ok && len(results) == 2 {
		if cleanedInt, ok := results[0].(int64); ok {
			cleaned = int(cleanedInt)
		}
		if returnedStr, ok := results[1].(string); ok {
			_, _ = fmt.Sscanf(returnedStr, "%f", &returned)
		}
	}

	return cleaned, returned, nil
}
