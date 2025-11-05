package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	InitRedis(context.Background())
	// 載入 Risk Pool 測試所需的 Lua 腳本
	InitLuaScripts(context.Background())
}

// TestRiskPoolReserve_Success 測試成功的預扣
func TestRiskPoolReserve_Success(t *testing.T) {
	ctx := context.Background()
	gameName := "BLACK_JACK"
	currency := "CNY_TEST"
	betID := "TEST_BET_001"

	// 清理測試數據
	defer cleanupTestData(t, gameName, currency)

	// 設置初始獎池
	setupTestPool(t, gameName, currency, 10000)

	// 執行 Reserve
	result, err := RiskPoolReserve(ctx, gameName, currency, betID, 120, 20*time.Minute, 10)

	// 驗證
	require.NoError(t, err)
	assert.Equal(t, 2, result.Status, "應該返回 reserved_ok 狀態")
	assert.Equal(t, 120.0, result.ReservedAmount, "預扣金額應為 120")

	// 驗證獎池金額
	poolAmount := getPoolAmount(t, gameName, currency)
	assert.Equal(t, 9880.0, poolAmount, "獎池應該減少 120")

	// 驗證預扣記錄
	resvAmt := getReservedAmount(t, gameName, currency, betID)
	assert.Equal(t, 120.0, resvAmt, "預扣記錄應該存在")
}

// TestRiskPoolReserve_Idempotent 測試冪等性
func TestRiskPoolReserve_Idempotent(t *testing.T) {
	ctx := context.Background()
	gameName := "BLACK_JACK"
	currency := "CNY_TEST"
	betID := "TEST_BET_002"

	defer cleanupTestData(t, gameName, currency)
	setupTestPool(t, gameName, currency, 10000)

	// 第一次 Reserve
	result1, err := RiskPoolReserve(ctx, gameName, currency, betID, 120, 20*time.Minute, 10)
	require.NoError(t, err)
	assert.Equal(t, 2, result1.Status)

	// 第二次 Reserve（相同 bet_id）
	result2, err := RiskPoolReserve(ctx, gameName, currency, betID, 120, 20*time.Minute, 10)
	require.NoError(t, err)
	assert.Equal(t, 1, result2.Status, "應該返回 already_reserved 狀態")
	assert.Equal(t, 120.0, result2.ReservedAmount)

	// 獎池金額應該保持不變
	poolAmount := getPoolAmount(t, gameName, currency)
	assert.Equal(t, 9880.0, poolAmount, "獎池不應該重複扣款")
}

// TestRiskPoolReserve_Insufficient 測試餘額不足
func TestRiskPoolReserve_Insufficient(t *testing.T) {
	ctx := context.Background()
	gameName := "BLACK_JACK"
	currency := "CNY_TEST"
	betID := "TEST_BET_003"

	defer cleanupTestData(t, gameName, currency)
	setupTestPool(t, gameName, currency, 100) // 只有 100

	// 嘗試 Reserve 120
	result, err := RiskPoolReserve(ctx, gameName, currency, betID, 120, 20*time.Minute, 10)

	require.NoError(t, err)
	assert.Equal(t, 0, result.Status, "應該返回 insufficient 狀態")
	assert.Equal(t, 100.0, result.AvailableNow, "應該顯示當前可用金額")

	// 獎池金額不變
	poolAmount := getPoolAmount(t, gameName, currency)
	assert.Equal(t, 100.0, poolAmount, "獎池金額不應該改變")
}

// TestRiskPoolSettle_Win 測試玩家贏的結算
func TestRiskPoolSettle_Win(t *testing.T) {
	ctx := context.Background()
	gameName := "BLACK_JACK"
	currency := "CNY_TEST"
	betID := "TEST_BET_004"

	defer cleanupTestData(t, gameName, currency)
	setupTestPool(t, gameName, currency, 10000)

	// 先 Reserve
	_, err := RiskPoolReserve(ctx, gameName, currency, betID, 120, 20*time.Minute, 10)
	require.NoError(t, err)

	// Settle Win
	result, err := RiskPoolSettle(ctx, gameName, currency, betID, "win", 78.8, 38)

	require.NoError(t, err)
	assert.Equal(t, 2, result.Status, "應該返回 success 狀態")
	assert.InDelta(t, 41.2, result.DeltaReturn, 0.01, "退回差額應為 41.2")
	assert.Equal(t, 38.0, result.HouseGainAdded, "加入投注額應為 38")
	assert.Equal(t, 120.0, result.OriginalReserved, "原預扣金額應為 120")

	// 驗證獎池金額
	// 9880 (reserve 後) + 41.2 (退回差額) + 38 (投注額) = 9959.2
	poolAmount := getPoolAmount(t, gameName, currency)
	assert.InDelta(t, 9959.2, poolAmount, 0.01, "獎池金額應為 9959.2")

	// 預扣應該被清理
	resvAmt := getReservedAmount(t, gameName, currency, betID)
	assert.Equal(t, 0.0, resvAmt, "預扣記錄應該被刪除")

	// 應該被標記為已結算
	isSettled := isSettled(t, gameName, currency, betID)
	assert.True(t, isSettled, "應該被標記為已結算")
}

// TestRiskPoolSettle_Lose 測試玩家輸的結算
func TestRiskPoolSettle_Lose(t *testing.T) {
	ctx := context.Background()
	gameName := "BLACK_JACK"
	currency := "CNY_TEST"
	betID := "TEST_BET_005"

	defer cleanupTestData(t, gameName, currency)
	setupTestPool(t, gameName, currency, 10000)

	// 先 Reserve
	_, err := RiskPoolReserve(ctx, gameName, currency, betID, 120, 20*time.Minute, 10)
	require.NoError(t, err)

	// Settle Lose
	result, err := RiskPoolSettle(ctx, gameName, currency, betID, "lose", 0, 38)

	require.NoError(t, err)
	assert.Equal(t, 2, result.Status, "應該返回 success 狀態")
	assert.Equal(t, 120.0, result.DeltaReturn, "退回金額應為 120")
	assert.Equal(t, 38.0, result.HouseGainAdded, "加入投注額應為 38")

	// 驗證獎池金額
	// 9880 (reserve 後) + 120 (退回) + 38 (投注額) = 10038
	poolAmount := getPoolAmount(t, gameName, currency)
	assert.Equal(t, 10038.0, poolAmount, "獎池金額應為 10038")
}

// TestRiskPoolSettle_Idempotent 測試結算冪等性
func TestRiskPoolSettle_Idempotent(t *testing.T) {
	ctx := context.Background()
	gameName := "BLACK_JACK"
	currency := "CNY_TEST"
	betID := "TEST_BET_006"

	defer cleanupTestData(t, gameName, currency)
	setupTestPool(t, gameName, currency, 10000)

	// Reserve + Settle
	_, err := RiskPoolReserve(ctx, gameName, currency, betID, 120, 20*time.Minute, 10)
	require.NoError(t, err)

	result1, err := RiskPoolSettle(ctx, gameName, currency, betID, "lose", 0, 38)
	require.NoError(t, err)
	assert.Equal(t, 2, result1.Status)

	poolAfterFirstSettle := getPoolAmount(t, gameName, currency)

	// 第二次 Settle（相同 bet_id）
	result2, err := RiskPoolSettle(ctx, gameName, currency, betID, "lose", 0, 38)
	require.NoError(t, err)
	assert.Equal(t, 3, result2.Status, "應該返回 already_settled 狀態")

	// 獎池金額應該保持不變
	poolAfterSecondSettle := getPoolAmount(t, gameName, currency)
	assert.Equal(t, poolAfterFirstSettle, poolAfterSecondSettle, "獎池不應該重複結算")
}

// TestRiskPoolGarbageCollection 測試垃圾回收機制
func TestRiskPoolGarbageCollection(t *testing.T) {
	ctx := context.Background()
	gameName := "BLACK_JACK"
	currency := "CNY_TEST"

	defer cleanupTestData(t, gameName, currency)
	setupTestPool(t, gameName, currency, 10000)

	// 創建一個已過期的預扣（TTL 設為 1 毫秒）
	betID1 := "TEST_BET_EXPIRED"
	_, err := RiskPoolReserve(ctx, gameName, currency, betID1, 100, 1*time.Millisecond, 10)
	require.NoError(t, err)

	// 等待過期
	time.Sleep(10 * time.Millisecond)

	// 創建新的預扣，應該觸發 GC
	betID2 := "TEST_BET_NEW"
	result, err := RiskPoolReserve(ctx, gameName, currency, betID2, 50, 20*time.Minute, 10)
	require.NoError(t, err)
	assert.Equal(t, 2, result.Status)
	assert.Equal(t, int64(1), result.GCCount, "應該清理了 1 個過期的預扣")

	// 驗證獎池金額
	// 10000 - 100 (expired) + 100 (GC returned) - 50 (new) = 9950
	poolAmount := getPoolAmount(t, gameName, currency)
	assert.Equal(t, 9950.0, poolAmount, "過期的預扣應該被回收")
}

// TestRiskPoolConcurrent 測試並發安全性
func TestRiskPoolConcurrent(t *testing.T) {
	ctx := context.Background()
	gameName := "BLACK_JACK"
	currency := "CNY_TEST"

	defer cleanupTestData(t, gameName, currency)
	setupTestPool(t, gameName, currency, 200) // 只有 200

	// 3 個並發請求，每個預扣 120
	done := make(chan *RiskPoolReserveResult, 3)
	for i := 0; i < 3; i++ {
		go func(idx int) {
			betID := fmt.Sprintf("CONCURRENT_BET_%d", idx)
			result, err := RiskPoolReserve(ctx, gameName, currency, betID, 120, 20*time.Minute, 10)
			if err != nil {
				t.Logf("Reserve error: %v", err)
			}
			done <- result
		}(i)
	}

	// 收集結果
	successCount := 0
	insufficientCount := 0
	for i := 0; i < 3; i++ {
		result := <-done
		if result != nil {
			if result.Status == 2 {
				successCount++
			} else if result.Status == 0 {
				insufficientCount++
			}
		}
	}

	// 應該只有 1 個成功，2 個失敗
	assert.Equal(t, 1, successCount, "應該只有 1 個請求成功")
	assert.Equal(t, 2, insufficientCount, "應該有 2 個請求因餘額不足失敗")

	// 獎池應該是 80 (200 - 120)
	poolAmount := getPoolAmount(t, gameName, currency)
	assert.Equal(t, 80.0, poolAmount, "獎池不應該超賣")
}

// TestRiskPoolSettle_WithoutReservation 測試沒有預扣的結算
func TestRiskPoolSettle_WithoutReservation(t *testing.T) {
	ctx := context.Background()
	gameName := "BLACK_JACK"
	currency := "CNY_TEST"
	betID := "TEST_BET_NO_RESERVE"

	defer cleanupTestData(t, gameName, currency)
	setupTestPool(t, gameName, currency, 10000)

	// 直接 Settle（沒有 Reserve）
	result, err := RiskPoolSettle(ctx, gameName, currency, betID, "lose", 0, 38)

	require.NoError(t, err)
	assert.Equal(t, 4, result.Status, "應該返回 settled_without_reservation 狀態")
	assert.Equal(t, 38.0, result.HouseGainAdded, "應該加入 house_gain")

	// 獎池應該增加 38
	poolAmount := getPoolAmount(t, gameName, currency)
	assert.Equal(t, 10038.0, poolAmount, "獎池應該增加 38")
}

// 輔助函數

func setupTestPool(t *testing.T, gameName, currency string, amount float64) {
	key := GetRiskPoolAvailableKey(gameName, currency)
	err := redisConn.Set(context.Background(), key, amount, 0).Err()
	require.NoError(t, err, "設置測試獎池失敗")
}

func getPoolAmount(t *testing.T, gameName, currency string) float64 {
	key := GetRiskPoolAvailableKey(gameName, currency)
	val, err := redisConn.Get(context.Background(), key).Float64()
	if err != nil {
		return 0
	}
	return val
}

func getReservedAmount(t *testing.T, gameName, currency, betID string) float64 {
	key := GetRiskPoolResvAmtKey(gameName, currency)
	val, err := redisConn.HGet(context.Background(), key, betID).Float64()
	if err != nil {
		return 0
	}
	return val
}

func isSettled(t *testing.T, gameName, currency, betID string) bool {
	key := GetRiskPoolSettledSetKey(gameName, currency)
	val, err := redisConn.SIsMember(context.Background(), key, betID).Result()
	if err != nil {
		return false
	}
	return val
}

func cleanupTestData(t *testing.T, gameName, currency string) {
	ctx := context.Background()
	keys := []string{
		GetRiskPoolAvailableKey(gameName, currency),
		GetRiskPoolResvAmtKey(gameName, currency),
		GetRiskPoolResvZKey(gameName, currency),
		GetRiskPoolSettledSetKey(gameName, currency),
	}
	for _, key := range keys {
		_ = redisConn.Del(ctx, key)
	}
}
