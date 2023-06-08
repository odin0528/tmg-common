package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
	"xxx/common/configs"
	"xxx/common/logs"
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

	if isLock := RedisLock(WOW_GAMING_MUTEX_PREFIX, key); !isLock {
		t.Fatal("Redis lock lock failed")
	}

	if isLock := RedisLock(WOW_GAMING_MUTEX_PREFIX, key); isLock {
		t.Fatal("Redis lock duplicate lock")
	}

	if isLock := RedisUnlock(WOW_GAMING_MUTEX_PREFIX, key); !isLock {
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

	fileds, err := GetHashMapFileds(key)
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

	fileds, err = GetHashMapFileds(key)
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
