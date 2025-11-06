package redis

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// Risk Pool Redis Keys
func GetRiskPoolAvailableKey(gameName, currency string) string {
	// 使用原本的 risk_prevent_pool:amount key，保持一致性
	gameName = strings.ToUpper(gameName)
	return fmt.Sprintf("risk_prevent_pool:amount:%s:%s", gameName, currency)
}

func GetRiskPoolResvAmtKey(gameName, currency string) string {
	gameName = strings.ToUpper(gameName)
	return fmt.Sprintf("risk_prevent_pool:resv:amt:%s:%s", gameName, currency)
}

func GetRiskPoolResvZKey(gameName, currency string) string {
	gameName = strings.ToUpper(gameName)
	return fmt.Sprintf("risk_prevent_pool:resv:exp:%s:%s", gameName, currency)
}

func GetRiskPoolSettledSetKey(gameName, currency string) string {
	gameName = strings.ToUpper(gameName)
	return fmt.Sprintf("risk_prevent_pool:settled:%s:%s", gameName, currency)
}

// RiskPoolReserveResult represents the result of reserve operation
type RiskPoolReserveResult struct {
	Status         int // 0=insufficient, 1=already_reserved, 2=reserved_ok
	ReservedAmount float64
	AvailableNow   float64 // for status=0, shows current available
	GCCount        int64   // how many expired reservations were cleaned
}

// RiskPoolSettleResult represents the result of settle operation
type RiskPoolSettleResult struct {
	Status           int     // 0=error, 2=success, 3=already_settled, 4=settled_without_reservation
	DeltaReturn      float64 // how much was returned to pool
	HouseGainAdded   float64 // how much house_gain was added
	OriginalReserved float64 // original reserved amount
}

// RiskPoolReserve executes the reserve Lua script
// Returns: result, error
func RiskPoolReserve(ctx context.Context, gameName, currency, betID string, amount float64, ttl time.Duration, gcLimit int64) (*RiskPoolReserveResult, error) {
	availableKey := GetRiskPoolAvailableKey(gameName, currency)
	resvAmtKey := GetRiskPoolResvAmtKey(gameName, currency)
	resvZKey := GetRiskPoolResvZKey(gameName, currency)

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

	result := &RiskPoolReserveResult{
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

// RiskPoolSettle executes the settle Lua script
func RiskPoolSettle(ctx context.Context, gameName, currency, betID, result string, actualPayout, houseGain float64) (*RiskPoolSettleResult, error) {
	availableKey := GetRiskPoolAvailableKey(gameName, currency)
	resvAmtKey := GetRiskPoolResvAmtKey(gameName, currency)
	resvZKey := GetRiskPoolResvZKey(gameName, currency)
	settledSetKey := GetRiskPoolSettledSetKey(gameName, currency)

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

	return &RiskPoolSettleResult{
		Status:           int(status),
		DeltaReturn:      deltaReturn,
		HouseGainAdded:   gainAdded,
		OriginalReserved: originalReserved,
	}, nil
}

// GetRiskPoolAvailable gets the current available pool amount
func GetRiskPoolAvailable(ctx context.Context, gameName, currency string) (float64, error) {
	key := GetRiskPoolAvailableKey(gameName, currency)
	val, err := redisConn.Get(ctx, key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return 0, nil
		}
		return 0, err
	}
	return parseFloatString(val)
}

// SetRiskPoolAvailable sets the available pool amount (for initialization/admin)
func SetRiskPoolAvailable(ctx context.Context, gameName, currency string, amount float64) error {
	key := GetRiskPoolAvailableKey(gameName, currency)
	return redisConn.Set(ctx, key, amount, 0).Err()
}

// Helper function to parse float from Redis result
func parseFloat(v interface{}) (float64, bool) {
	switch val := v.(type) {
	case int64:
		return float64(val), true
	case float64:
		return val, true
	case string:
		f, err := parseFloatString(val)
		return f, err == nil
	default:
		return 0, false
	}
}

func parseFloatString(s string) (float64, error) {
	var f float64
	_, err := fmt.Sscanf(s, "%f", &f)
	return f, err
}

// CleanupExpiredReserves 清理過期預扣
func CleanupExpiredReserves(ctx context.Context, gameName, currency string, limit int) (cleaned int, returned float64, err error) {
	availableKey := GetRiskPoolAvailableKey(gameName, currency)
	resvAmtKey := GetRiskPoolResvAmtKey(gameName, currency)
	resvZKey := GetRiskPoolResvZKey(gameName, currency)

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
