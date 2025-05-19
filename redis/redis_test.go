package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"game_server/common/configs"
	"game_server/common/logs"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/stretchr/testify/assert"
)

func init() {
	initConfig()
	logs.InitLogs()
}

func initConfig() {
	configs.Init("../configs/example/")
	configs.LoadConfig([]string{"../configs/example/cache.ini", "../configs/example/env.ini"})
}

func TestAll(t *testing.T) {
	TestRedisConnect(t)
	TestRedisString(t)
	TestRedisInt(t)
	TestRedisFloat(t)
	TestRedisBool(t)
	TestRedisStruct(t)
	TestIncrease(t)
}

func TestRedisConnect(t *testing.T) {
	err := InitRedis(context.Background())
	if err != nil {
		t.Fatal("InitRedis err:", err.Error())
	}

	t.Log("Connect Redis ok")
}

func TestRedisString(t *testing.T) {
	err := InitRedis(context.Background())
	if err != nil {
		t.Fatal("InitRedis err:", err.Error())
	}

	key := "testKey"
	value := "testValue"
	if err := Put(key, value, 0); err != nil {
		fmt.Println("Put key :", key, "err=", err.Error())
		return
	}

	if retValue, ok := GetString(key); !ok {
		t.Fatal("Get string value failed")
	} else if retValue != value {
		t.Fatal(fmt.Sprintf("Get string value is not expect. real value :%s , expect value :%s", retValue, value))
		fmt.Println("Get value =", retValue)
	}

	if err := Delete([]string{key}); err != nil {
		t.Fatal("Delete cache key failed")
	}

	t.Log("Redis string ok")
}

func TestRedisInt(t *testing.T) {
	err := InitRedis(context.Background())
	if err != nil {
		t.Fatal("InitRedis err:", err.Error())
	}

	key := "testKey"
	value := 99
	if err := Put(key, value, 0); err != nil {
		fmt.Println("Put key :", key, "err=", err.Error())
		return
	}

	if retValue, ok := GetInt(key); !ok {
		t.Fatal("Get int value failed")
	} else if retValue != value {
		t.Fatal(fmt.Sprintf("Get int value is not expect. real value :%v , expect value :%v", retValue, value))
		fmt.Println("Get value =", retValue)
	}

	if err := Delete([]string{key}); err != nil {
		t.Fatal("Delete cache key failed")
	}

	t.Log("Redis int ok")
}

func TestRedisFloat(t *testing.T) {
	err := InitRedis(context.Background())
	if err != nil {
		t.Fatal("InitRedis err:", err.Error())
	}

	key := "testKey"
	value := 99.9
	if err := Put(key, value, 0); err != nil {
		fmt.Println("Put key :", key, "err=", err.Error())
		return
	}

	if retValue, ok := GetFloat64(key); !ok {
		t.Fatal("Get float value failed")
	} else if retValue != value {
		t.Fatal(fmt.Sprintf("Get float value is not expect. real value :%v , expect value :%v", retValue, value))
		fmt.Println("Get value =", retValue)
	}

	if err := Delete([]string{key}); err != nil {
		t.Fatal("Delete cache key failed")
	}

	t.Log("Redis float ok")
}

func TestRedisBool(t *testing.T) {
	err := InitRedis(context.Background())
	if err != nil {
		t.Fatal("InitRedis err:", err.Error())
	}

	key := "testKey"
	value := true
	if err := Put(key, value, 0); err != nil {
		fmt.Println("Put key :", key, "err=", err.Error())
		return
	}

	if retValue, ok := GetBool(key); !ok {
		t.Fatal("Get bool value failed")
	} else if retValue != value {
		t.Fatal(fmt.Sprintf("Get bool value is not expect. real value :%v , expect value :%v", retValue, value))
		fmt.Println("Get value =", retValue)
	}

	if err := Delete([]string{key}); err != nil {
		t.Fatal("Delete cache key failed")
	}

	t.Log("Redis bool ok")
}

func TestRedisStruct(t *testing.T) {
	type redisStore struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	}

	err := InitRedis(context.Background())
	if err != nil {
		t.Fatal("InitRedis err:", err.Error())
	}

	key := "testKey"
	value := redisStore{
		Name: "test_player",
		Id:   9999,
	}

	if err := Put(key, value, 0); err != nil {
		fmt.Println("Put key :", key, "err=", err.Error())
		return
	}

	retValue := redisStore{}
	if ok := GetStructData(key, &retValue); !ok {
		t.Fatal("Get struct value failed")
	} else if retValue.Id != value.Id || retValue.Name != value.Name {
		t.Fatal(fmt.Sprintf("Get struct id is not expect. real value :%v , expect value :%v", retValue, value))
	}

	if err := Delete([]string{key}); err != nil {
		t.Fatal("Delete cache key failed")
	}

	t.Log("Redis struct ok")
}

func TestRedisLock(t *testing.T) {
	err := InitRedis(context.Background())
	if err != nil {
		t.Fatal("InitRedis err:", err.Error())
	}

	key := "lock_key"

	if isLock := RedisLock(GetCacheKey(WOW_GAMING_MUTEX_PREFIX, key)); !isLock {
		t.Fatal("Redis lock lock failed")
	}

	if isLock := RedisLock(GetCacheKey(WOW_GAMING_MUTEX_PREFIX, key)); isLock {
		t.Fatal("Redis lock duplicate lock")
	}

	if isLock := RedisUnlock(GetCacheKey(WOW_GAMING_MUTEX_PREFIX, key)); !isLock {
		t.Fatal("Redis lock  unlock failed")
	}
}

func TestExist(t *testing.T) {
	InitRedis(context.Background())

	key := "key_test_exist"

	if IsExist(key) {
		info, ok := GetString(key)
		if !ok {
			t.Fatal("get redis failed")
		} else {
			t.Fatal("info:", info)
		}
		t.Fatal("Key can't exist")
	}

	if err := Put(key, 1, time.Second); err != nil {
		t.Fatal("Put key failed. err:", err.Error())
	}

	if !IsExist(key) {
		t.Fatal("Key can't find")
	}

	t.Log("Exist finish")
}

func TestHashMap(t *testing.T) {
	InitRedis(context.Background())
	key := "testHashKey"

	fileds, err := GetHashMapFields(key)
	if err != nil {
		t.Fatal("GetHashMapKeys err:", err.Error())
	}

	if len(fileds) > 0 {
		if err := DelHashMap(key, fileds); err != nil {
			t.Fatal("Del hash map faild. err:", err.Error())
		}
	}

	results, err := GetAllHashMap(key)
	if err != nil {
		t.Fatal("Redis get hash failed. er:", err.Error())
	} else if len(results) > 0 {
		t.Fatal("Put hash data error. real data=", results)
	}

	type MoneyInfo struct {
		Account string  `json:"acoount"`
		Money   float64 `json:"money"`
	}

	account := "player_1"
	account2 := "player_2"

	value := map[string]interface{}{}

	value[account] = 10.0
	value[account] = 20
	value[account2] = 50.5

	if err := PutInHashMap(key, value); err != nil {
		t.Fatal("Redis put hash failed. er:", err.Error())
	}

	moneyInfo := MoneyInfo{
		Account: "bot",
		Money:   1000.0,
	}

	value = map[string]interface{}{}
	value["bot"] = moneyInfo
	if err := PutInHashMap(key, value); err != nil {
		t.Fatal("Redis put hash failed. er:", err.Error())
	}

	results, err = GetAllHashMap(key)
	if err != nil {
		t.Fatal("Redis get hash failed. er:", err.Error())
	} else if len(results) != 3 {
		t.Fatal("Put hash data loss. real len=", len(results))
	}

	for k, v := range results {
		if k == "player_1" {
			num, err := strconv.Atoi(v)
			if err != nil {
				t.Fatal("atoi err:", err.Error())
			}

			if num != 20 {
				t.Log("get int failed")
			}
		} else if k == "player_2" {
			num, err := strconv.ParseFloat(v, 32)
			if err != nil {
				t.Fatal("atoi err:", err.Error())
			}

			if num != 50.5 {
				t.Log("get float failed")
			}
		} else if k == "bot" {
			info := MoneyInfo{}
			if err := json.Unmarshal([]byte(v), &info); err != nil {
				t.Fatal("Unmarshal money info failed. err:", err.Error())
			}
		}
	}

	if value, err := GetHashMap(key, "player_1"); err != nil {
		t.Fatal("Redis get hash failed. er:", err.Error())
	} else {
		num, err := strconv.Atoi(value)
		if err != nil {
			t.Fatal("atoi err:", err.Error())
		}

		if num != 20 {
			t.Fatal("get int failed")
		}
	}

	fileds, err = GetHashMapFields(key)
	if err != nil {
		t.Fatal("GetHashMapKeys err:", err.Error())
	}

	if err := DelHashMap(key, fileds); err != nil {
		t.Fatal("Del hash map faild. err:", err.Error())
	}
}

func TestHashMapReuse(t *testing.T) {
	InitRedis(context.Background())
	key := "testHashKey"
	account := "player_1"
	betIds := []string{"abc123", "zzz321"}
	value := map[string]interface{}{
		account: betIds,
	}

	if err := PutInHashMap(key, value); err != nil {
		t.Fatal("Redis put hash failed. er:", err.Error())
	}

	results, err := GetHashMap(key, account)
	if err != nil {
		t.Fatal("Redis get hash failed. er:", err.Error())
	}

	getList := []string{}
	if err := json.Unmarshal([]byte(results), &getList); err != nil {
		t.Fatal("Unmarshal err:", err.Error())
	} else if len(getList) != 2 {
		t.Fatal("data loss. getList:", getList)
	} else if getList[0] != betIds[0] || getList[1] != betIds[1] {
		t.Fatal("data is not equal. get list:", getList, "origin list:", betIds)
	}

	betIds = append(betIds, "qqq132")
	value[account] = betIds
	if err := PutInHashMap(key, value); err != nil {
		t.Fatal("Redis put hash failed. er:", err.Error())
	}

	results, err = GetHashMap(key, account)
	if err != nil {
		t.Fatal("Redis get hash failed. er:", err.Error())
	}

	getList = []string{}
	if err := json.Unmarshal([]byte(results), &getList); err != nil {
		t.Fatal("Unmarshal err:", err.Error())
	} else if len(getList) != 3 {
		t.Fatal("data loss. getList:", getList)
	} else if getList[0] != betIds[0] || getList[1] != betIds[1] || getList[2] != betIds[2] {
		t.Fatal("data is not equal. get list:", getList, "origin list:", getList)
	}

	t.Log("getList:", getList)
	if err := DelHashMap(key, []string{account}); err != nil {
		t.Fatal("Del hash map faild. err:", err.Error())
	}
}

func TestIncreaseHashMap(t *testing.T) {
	InitRedis(context.Background())
	hashKey := "incrHashKey"
	filed := "player_1"
	filed2 := "player_2"
	num := 1

	if err := DelHashMap(hashKey, []string{filed, filed2}); err != nil {
		t.Fatal("Delete hash key err:", err.Error())
	}

	if count, err := IncreaseInHashMap(hashKey, filed, num); err != nil {
		t.Fatal("IncreaseInHashMap err:", err.Error())
	} else {
		t.Log("count:", count)
	}

	if resultMap, err := GetAllHashMap(hashKey); err != nil {
		t.Fatal("GetAllHashMap err:", err.Error())
	} else if len(resultMap) != 1 {
		t.Fatal("GetAllHashMap data loss:", resultMap)
	} else {
		t.Log("resultMap:", resultMap)
	}

	num = 3
	if count, err := IncreaseInHashMap(hashKey, filed, num); err != nil {
		t.Fatal("IncreaseInHashMap err:", err.Error())
	} else {
		t.Log("count:", count)
	}

	if resultMap, err := GetAllHashMap(hashKey); err != nil {
		t.Fatal("GetAllHashMap err:", err.Error())
	} else if len(resultMap) != 1 {
		t.Fatal("GetAllHashMap data loss:", resultMap)
	} else {
		t.Log("resultMap:", resultMap)
	}

	if count, err := IncreaseInHashMap(hashKey, filed2, num); err != nil {
		t.Fatal("IncreaseInHashMap err:", err.Error())
	} else {
		t.Log("count:", count)
	}

	if resultMap, err := GetAllHashMap(hashKey); err != nil {
		t.Fatal("GetAllHashMap err:", err.Error())
	} else if len(resultMap) != 2 {
		t.Fatal("GetAllHashMap data loss:", resultMap)
	} else {
		t.Log("resultMap:", resultMap)
	}

	num = -2
	if count, err := IncreaseInHashMap(hashKey, filed2, num); err != nil {
		t.Fatal("IncreaseInHashMap err:", err.Error())
	} else {
		t.Log("count:", count)
	}

	if resultMap, err := GetAllHashMap(hashKey); err != nil {
		t.Fatal("GetAllHashMap err:", err.Error())
	} else if len(resultMap) != 2 {
		t.Fatal("GetAllHashMap data loss:", resultMap)
	} else {
		t.Log("resultMap:", resultMap)
	}

	if err := DelHashMap(hashKey, []string{filed, filed2}); err != nil {
		t.Fatal("Delete hash key err:", err.Error())
	}
}

func TestPushAndPop(t *testing.T) {
	InitRedis(context.Background())
	testTime := 3
	key := "push_pop_key"

	for i := 0; i < testTime; i++ {
		if err := LPush(key, i); err != nil {
			t.Fatal("LPush err:", err.Error())
		}
	}

	for i := 0; i < testTime; i++ {
		if valueStr, ok := RPop(key); !ok {
			t.Fatal("RPop failed")
		} else if value, err := strconv.Atoi(valueStr); err != nil {
			t.Fatal("Convert value to int failed. err:", err.Error())
		} else if i != value {
			t.Fatal("value is err.", value)
		} else {
			t.Log("value =", value)
		}
	}
}

func TestIncrease(t *testing.T) {
	InitRedis(context.Background())
	key := "incKey"

	if err := Increase(key); err != nil {
		t.Fatal("Increase key err:", err.Error())
	}

	value, ok := GetInt(key)
	if !ok {
		t.Fatal("GetInt failed")
	}

	t.Log("value:", value)
}

func TestRedisSet(t *testing.T) {
	err := InitRedis(context.Background())
	if err != nil {
		t.Fatal("InitRedis err:", err.Error())
	}

	key := "testKey"

	data := []string{}
	for i := 0; i < 50000; i++ {
		data = append(data, fmt.Sprintf("value_%d", i%10000)) // 刻意加入重複元素
	}
	if err := SAdd(key, data); err != nil {
		t.Fatal("SAdd err:", err.Error())
	}

	count, err := SCard(key)
	assert.Nil(t, err)
	assert.Equal(t, count, int64(10000))

	// 批處理
	cnt := 0
	err = SScanWithCallback(key, 1000, func(members []string) error {
		// do nothing
		cnt += len(members)
		return nil
	})
	assert.Nil(t, err)
	assert.Equal(t, cnt, 10000)

	// 取出1000個元素
	values, err := SPopN(key, 1000)
	assert.Nil(t, err)
	assert.Equal(t, len(values), 1000)

	count, err = SCard(key)
	assert.Nil(t, err)
	assert.Equal(t, count, int64(9000))

	// 取出1個元素
	_, err = SPop(key)
	assert.Nil(t, err)

	count, err = SCard(key)
	assert.Nil(t, err)
	assert.Equal(t, count, int64(8999))

	if err := Delete([]string{key}); err != nil {
		t.Fatal("Delete cache key failed")
	}

	t.Log("Redis set ok")
}

func TestZIncrBy(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "localhost", "6379"),
		Password: "123456",
		DB:       0,
		PoolSize: 100,
	})
	SetRedisClient(redisClient)
	key := "rankings"

	_, err := ZIncrBy(key, 1, "TITAN_MONSTER")
	assert.Nil(t, err)

	_, err = ZIncrBy(key, 2, "ALICE_AND_WONDERLAND")
	assert.Nil(t, err)

	_, err = ZIncrBy(key, 3, "CALL_UFO")
	assert.Nil(t, err)

	results, err := ZRangeWithScores(key, 0, -1)
	assert.Nil(t, err)

	assert.Equal(t, len(results), 3)
	assert.Equal(t, results[0].Member, "TITAN_MONSTER")
	assert.Equal(t, results[1].Member, "ALICE_AND_WONDERLAND")
	assert.Equal(t, results[2].Member, "CALL_UFO")

	results, err = ZRevRangeWithScores(key, 0, -1)
	assert.Nil(t, err)

	assert.Equal(t, len(results), 3)
	assert.Equal(t, results[0].Member, "CALL_UFO")
	assert.Equal(t, results[1].Member, "ALICE_AND_WONDERLAND")
	assert.Equal(t, results[2].Member, "TITAN_MONSTER")

	Delete([]string{key})
}

func TestPipelined(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "localhost", "6379"),
		Password: "123456",
		DB:       0,
		PoolSize: 100,
	})
	SetRedisClient(redisClient)

	keysCount := 1000

	results, err := RunPipelined(context.Background(), func(p Pipeline) {
		for i := 0; i < keysCount; i++ {
			key := fmt.Sprintf("key_%d", i)
			p.Set(key, i, time.Minute)
		}
	})
	assert.Nil(t, err)

	for _, res := range results {
		assert.Nil(t, res.Err)
	}

	results, err = RunPipelined(context.Background(), func(p Pipeline) {
		for i := 0; i < keysCount; i++ {
			key := fmt.Sprintf("key_%d", i)
			p.Get(key)
		}
	})

	for i, res := range results {
		assert.Equal(t, res.Result, fmt.Sprintf("%d", i))
		assert.Nil(t, res.Err)
	}
}

func TestPipeline(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "localhost", "6379"),
		Password: "123456",
		DB:       0,
		PoolSize: 100,
	})
	SetRedisClient(redisClient)

	keysCount := 1000

	pipe := GetPipeline()

	for i := 0; i < keysCount; i++ {
		key := fmt.Sprintf("key_%d", i)
		pipe.Set(key, i, time.Minute)
	}
	_, err := pipe.Exec(context.Background())
	assert.Nil(t, err)

	pipe = GetPipeline()
	for i := 0; i < keysCount; i++ {
		key := fmt.Sprintf("key_%d", i)
		pipe.Get(key)
	}
	commands, err := pipe.Exec(context.Background())
	assert.Nil(t, err)

	for i, res := range commands {
		assert.Equal(t, res.Result, fmt.Sprintf("%d", i))
		assert.Nil(t, res.Err)
	}
}

func TestPipelineWithContext(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "localhost", "6379"),
		Password: "123456",
		DB:       0,
		PoolSize: 100,
	})
	SetRedisClient(redisClient)

	keysCount := 1000

	ctx := context.Background()

	ctx = WithPipeline(ctx, GetPipeline())

	pipe, ok := GetPipelineWithContext(ctx)
	assert.True(t, ok)

	for i := 0; i < keysCount; i++ {
		key := fmt.Sprintf("key_%d", i)
		pipe.Set(key, i, time.Minute)
	}
	_, err := pipe.Exec(context.Background())
	assert.Nil(t, err)

	pipe = GetPipeline()
	for i := 0; i < keysCount; i++ {
		key := fmt.Sprintf("key_%d", i)
		pipe.Get(key)
	}
	commands, err := pipe.Exec(context.Background())
	assert.Nil(t, err)

	for i, res := range commands {
		assert.Equal(t, res.Result, fmt.Sprintf("%d", i))
		assert.Nil(t, res.Err)
	}
}

func TestTxPipeline(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "localhost", "6379"),
		Password: "123456",
		DB:       0,
		PoolSize: 100,
	})
	SetRedisClient(redisClient)

	keysCount := 1000

	pipe := GetTxPipeline()

	for i := 0; i < keysCount; i++ {
		key := fmt.Sprintf("key_%d", i)
		pipe.Set(key, i, time.Minute)
	}
	_, err := pipe.Exec(context.Background())
	assert.Nil(t, err)

	pipe = GetTxPipeline()
	for i := 0; i < keysCount; i++ {
		key := fmt.Sprintf("key_%d", i)
		pipe.Get(key)
	}
	commands, err := pipe.Exec(context.Background())
	assert.Nil(t, err)

	for i, res := range commands {
		assert.Equal(t, res.Result, fmt.Sprintf("%d", i))
		assert.Nil(t, res.Err)
	}
}

func TestTxPipelineWithContext(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "localhost", "6379"),
		Password: "123456",
		DB:       0,
		PoolSize: 100,
	})
	SetRedisClient(redisClient)

	keysCount := 10

	ctx := context.Background()

	ctx = WithTxPipeline(ctx, GetTxPipeline())

	pipe, ok := GetTxPipelineWithContext(ctx)
	assert.True(t, ok)

	for i := 0; i < keysCount; i++ {
		key := fmt.Sprintf("key_%d", i)
		pipe.Set(key, i, time.Minute)
	}

	_, err := pipe.Exec(context.Background())
	assert.Nil(t, err)

	pipe, ok = GetTxPipelineWithContext(ctx)
	assert.True(t, ok)
	for i := 0; i < keysCount; i++ {
		key := fmt.Sprintf("key_%d", i)
		pipe.Get(key)
	}
	commands, err := pipe.Exec(context.Background())
	assert.Nil(t, err)

	for i, res := range commands {
		assert.Equal(t, res.Result, fmt.Sprintf("%d", i))
		assert.Nil(t, res.Err)
	}
}

func TestTxPipelineWithContext_Discard(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "localhost", "6379"),
		Password: "123456",
		DB:       0,
		PoolSize: 100,
	})
	SetRedisClient(redisClient)

	keysCount := 10

	ctx := context.Background()

	ctx = WithTxPipeline(ctx, GetTxPipeline())

	pipe, ok := GetTxPipelineWithContext(ctx)
	assert.True(t, ok)

	for i := 0; i < keysCount; i++ {
		key := fmt.Sprintf("key_%d", i)
		pipe.Set(key, i, time.Minute)
	}

	assert.Nil(t, DiscardTxPipeline(ctx))
	//_, err := pipe.Exec(context.Background()) // 因為已經執行 Discard 了，所以這次 Exec 將不會執行任何指令
	//assert.Nil(t, err)

	pipe, ok = GetTxPipelineWithContext(ctx)
	assert.True(t, ok)
	for i := 0; i < keysCount; i++ {
		key := fmt.Sprintf("key_%d", i)
		pipe.Get(key)
	}
	commands, err := pipe.Exec(context.Background())
	assert.Nil(t, err)

	for i, res := range commands {
		assert.Equal(t, res.Result, fmt.Sprintf("%d", i))
		assert.Nil(t, res.Err)
	}
}
