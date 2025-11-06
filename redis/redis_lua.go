package redis

import (
	"context"
	"fmt"
)

var (
	accessLimitScriptSHA string // 訪問限制

	riskPoolReserveSHA string // 水池預扣
	riskPoolSettleSHA  string // 水池結算
	riskPoolCleanupSHA string // 水池過期預扣清理
)

// Access Limit Lua 腳本
const accessLimitScriptContent = `
local count = redis.call('INCR', KEYS[1])
if count == 1 then
	redis.call('EXPIRE', KEYS[1], ARGV[1])
end
return count`

// Risk Pool Reserve Lua Script
// 預扣腳本：清理過期 + 檢查餘額 + 扣除可用額 + 記錄預扣
const riskPoolReserveScript = `
-- KEYS: [1]availableKey, [2]resvAmtKey, [3]resvZKey
-- ARGV: [1]bet_id, [2]amount, [3]now_ms, [4]expire_at_ms, [5]gc_limit

local availableKey = KEYS[1]
local resvAmtKey   = KEYS[2]
local resvZKey     = KEYS[3]

local bet_id       = ARGV[1]
local amount       = tonumber(ARGV[2])
local now_ms       = tonumber(ARGV[3])
local expire_at_ms = tonumber(ARGV[4])
local gc_limit     = tonumber(ARGV[5])

-- GC: Clean expired reservations (small batch)
local expired = redis.call('ZRANGEBYSCORE', resvZKey, 0, now_ms, 'LIMIT', 0, gc_limit)
local gc_count = 0
for i, bid in ipairs(expired) do
  local amt = tonumber(redis.call('HGET', resvAmtKey, bid) or '0')
  if amt > 0 then
    redis.call('INCRBYFLOAT', availableKey, amt)
    gc_count = gc_count + 1
  end
  redis.call('HDEL', resvAmtKey, bid)
  redis.call('ZREM', resvZKey, bid)
end

-- Idempotency: already reserved?
if redis.call('HEXISTS', resvAmtKey, bet_id) == 1 then
  local existing = tonumber(redis.call('HGET', resvAmtKey, bet_id))
  return {1, tostring(existing), gc_count} -- status=1 (already reserved), reserved_amount, gc_count
end

-- Check available pool
local available = tonumber(redis.call('GET', availableKey) or '0')
if available < amount then
  return {0, tostring(available), gc_count} -- status=0 (insufficient), current_available, gc_count
end

-- Reserve: deduct from available pool
redis.call('INCRBYFLOAT', availableKey, -amount)
redis.call('HSET', resvAmtKey, bet_id, tostring(amount))
redis.call('ZADD', resvZKey, expire_at_ms, bet_id)

return {2, tostring(amount), gc_count} -- status=2 (reserved ok), reserved_amount, gc_count
`

// Risk Pool Settle Lua Script
// 結算腳本：檢查冪等 + 處理輸贏 + 清理預扣 + 標記已結算
const riskPoolSettleScript = `
-- KEYS: [1]availableKey, [2]resvAmtKey, [3]resvZKey, [4]settledSetKey
-- ARGV: [1]bet_id, [2]result("win"|"lose"), [3]actual_payout, [4]house_gain

local availableKey  = KEYS[1]
local resvAmtKey    = KEYS[2]
local resvZKey      = KEYS[3]
local settledSetKey = KEYS[4]

local bet_id        = ARGV[1]
local result        = ARGV[2]
local actual_payout = tonumber(ARGV[3])
local house_gain    = tonumber(ARGV[4])

-- Idempotency: already settled?
if redis.call('SISMEMBER', settledSetKey, bet_id) == 1 then
  return {3, "0", "0", "0"} -- status=3 (already settled), delta_return, house_gain_added, 0
end

-- Get reserved amount
local reserved = tonumber(redis.call('HGET', resvAmtKey, bet_id) or '0')

local delta_return = 0
local gain_added = 0

if reserved == 0 then
  -- No reservation (might be expired and GC'd)
  if result == 'lose' and house_gain > 0 then
    -- Player loses: add house_gain to pool
    redis.call('INCRBYFLOAT', availableKey, house_gain)
    gain_added = house_gain
  elseif result == 'win' or result == 'tie' then
    -- Player wins or tie: add house_gain but deduct actual_payout
    local delta = house_gain - actual_payout
    redis.call('INCRBYFLOAT', availableKey, delta)
    gain_added = house_gain
  end
  redis.call('SADD', settledSetKey, bet_id)
  redis.call('EXPIRE', settledSetKey, 86400) -- expire settled set after 24h
  return {4, "0", tostring(gain_added), "0"} -- status=4 (settled without reservation)
end

-- Process win/lose/tie
if result == 'win' or result == 'tie' then
  -- Player wins or tie: return difference (reserved - actual_payout) + add house_gain
  if reserved > actual_payout then
    delta_return = reserved - actual_payout
  else
    delta_return = 0
  end
  
  -- Add house_gain (bet amount after levy) to pool
  redis.call('INCRBYFLOAT', availableKey, delta_return + house_gain)
  gain_added = house_gain
  
  -- If reserved < actual_payout (shouldn't happen with correct max_liability)
  -- Pool might go negative (need monitoring)
  
elseif result == 'lose' then
  -- Player loses: return all reserved + add house_gain
  delta_return = reserved
  redis.call('INCRBYFLOAT', availableKey, reserved + house_gain)
  gain_added = house_gain
  
else
  -- Invalid result
  return {0, "0", "0", "0"} -- error
end

-- Clean up reservation
redis.call('HDEL', resvAmtKey, bet_id)
redis.call('ZREM', resvZKey, bet_id)
redis.call('SADD', settledSetKey, bet_id)
redis.call('EXPIRE', settledSetKey, 86400) -- expire after 24h

-- 使用 tostring 確保返回字符串格式的數字，保留小數點
return {2, tostring(delta_return), tostring(gain_added), tostring(reserved)} -- status=2 (success), delta_return, gain_added, original_reserved
`

// Risk Pool Cleanup Lua Script
// 清理過期預扣腳本：查找過期 + 歸還金額 + 刪除記錄
const riskPoolCleanupScript = `
-- KEYS: [1]availableKey, [2]resvAmtKey, [3]resvZKey
-- ARGV: [1]now_ms, [2]limit

local availableKey = KEYS[1]
local resvAmtKey   = KEYS[2]
local resvZKey     = KEYS[3]

local now_ms = tonumber(ARGV[1])
local limit  = tonumber(ARGV[2])

-- 查找過期的預扣
local expired = redis.call('ZRANGEBYSCORE', resvZKey, 0, now_ms, 'LIMIT', 0, limit)
local cleaned_count = 0
local total_returned = 0

for i, bet_id in ipairs(expired) do
  local amt = tonumber(redis.call('HGET', resvAmtKey, bet_id) or '0')
  if amt > 0 then
    redis.call('INCRBYFLOAT', availableKey, amt)
    total_returned = total_returned + amt
  end
  redis.call('HDEL', resvAmtKey, bet_id)
  redis.call('ZREM', resvZKey, bet_id)
  cleaned_count = cleaned_count + 1
end

return {cleaned_count, tostring(total_returned)}
`

// InitLuaScripts loads Lua scripts into Redis
func InitLuaScripts(ctx context.Context) error {
	// 載入 Access Limit Lua 腳本
	sha, err := redisConn.ScriptLoad(ctx, accessLimitScriptContent).Result()
	if err != nil {
		return fmt.Errorf("failed to load access limit script: %w", err)
	}
	accessLimitScriptSHA = sha

	// Load Reserve Script
	sha, err = redisConn.ScriptLoad(ctx, riskPoolReserveScript).Result()
	if err != nil {
		return fmt.Errorf("failed to load reserve script: %w", err)
	}
	riskPoolReserveSHA = sha

	// Load Settle Script
	sha, err = redisConn.ScriptLoad(ctx, riskPoolSettleScript).Result()
	if err != nil {
		return fmt.Errorf("failed to load settle script: %w", err)
	}
	riskPoolSettleSHA = sha

	// Load Cleanup Script
	sha, err = redisConn.ScriptLoad(ctx, riskPoolCleanupScript).Result()
	if err != nil {
		return fmt.Errorf("failed to load cleanup script: %w", err)
	}
	riskPoolCleanupSHA = sha

	return nil
}

// GetRiskPoolReserveSHA returns the SHA of reserve script
func GetRiskPoolReserveSHA() string {
	return riskPoolReserveSHA
}

// GetRiskPoolSettleSHA returns the SHA of settle script
func GetRiskPoolSettleSHA() string {
	return riskPoolSettleSHA
}

// GetRiskPoolCleanupSHA returns the SHA of cleanup script
func GetRiskPoolCleanupSHA() string {
	return riskPoolCleanupSHA
}
