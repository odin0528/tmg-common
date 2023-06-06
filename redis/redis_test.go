package redis

import (
	"context"
	"fmt"
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
