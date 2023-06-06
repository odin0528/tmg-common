package redis

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	"xxx/common/configs"
	"xxx/common/logs"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	jsoniter "github.com/json-iterator/go"
)

var redisConn *redis.Client
var mutexMap sync.Map

func InitRedis(ctx context.Context) error {
	section := configs.SECTION_CACHE

	host := configs.Get(section, configs.CACHE_HOST, "localhost")
	port := configs.Get(section, configs.CACHE_PORT, "6379")
	password := configs.Get(section, configs.CACHE_PASSWORD, "")
	dbNum := configs.GetInt(section, configs.CACHE_DB_NUM, 0)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       dbNum,
		PoolSize: 100,
	})

	if redisClient == nil {
		return errors.New("New redis client failed")
	}

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		// panic(err)
		return err
	}

	redisConn = redisClient

	return nil
}

func Put(key string, value interface{}, timeout time.Duration) (err error) {
	var putValue interface{}

	switch value.(type) {
	case string, bool, float32, float64, int, int8, int16, int32, int64:
		putValue = value
	default:
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		putValue, err = json.Marshal(value)

		if nil != err {
			logs.Error(
				logs.LOG_TYPE_SYSTEM,
				logs.LOG_KEY_CACHE,
				fmt.Sprintf("failed to Marshal value. err: %s", err.Error()),
				map[string]interface{}{
					logs.FIELD_KEY_CACHE_KEY: key,
					logs.FIELD_KEY_PAYLOAD:   value,
					logs.FIELD_KEY_TIMEOUT:   timeout,
				},
			)

			return err
		}
	}

	return redisConn.Set(context.Background(), key, putValue, timeout).Err()
}

func PutNoExpiry(key string, value interface{}) (err error) {
	var putValue interface{}

	switch value.(type) {
	case string, bool, float32, float64, int, int8, int16, int32, int64:
		putValue = value
	default:
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		putValue, err = json.Marshal(value)

		if nil != err {
			logs.Error(
				logs.LOG_TYPE_SYSTEM,
				logs.LOG_KEY_CACHE,
				fmt.Sprintf("failed to Marshal value. err: %s", err.Error()),
				map[string]interface{}{
					logs.FIELD_KEY_CACHE_KEY: key,
					logs.FIELD_KEY_PAYLOAD:   value,
				},
			)

			return err
		}
	}

	return redisConn.Set(context.Background(), key, putValue, 0).Err()
}

func Delete(keys []string) error {

	return redisConn.Del(context.Background(), keys...).Err()
}

func GetString(key string) (retValue string, ok bool) {
	value := redisConn.Get(context.Background(), key).Val()

	return value, true
}

func GetInt(key string) (retValue int, ok bool) {
	value, err := redisConn.Get(context.Background(), key).Int()
	if err != nil {
		return 0, false
	}

	return value, true
}

func GetFloat64(key string) (retValue float64, ok bool) {
	value, err := redisConn.Get(context.Background(), key).Float64()
	if err != nil {
		return 0, false
	}

	return value, true
}

func GetBool(key string) (retValue bool, ok bool) {
	byteValues, err := redisConn.Get(context.Background(), key).Bytes()
	if err != nil {
		return false, false
	}

	if len(byteValues) != 1 {
		return false, false
	}

	switch byteValues[0] {
	case 48:
		retValue = false
	case 49:
		retValue = true
	}

	return retValue, true
}

func GetStructData(key string, data interface{}) bool {
	value, ok := GetString(key)
	if ok {
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		err := json.Unmarshal(([]byte)(value), data)
		ok = (nil == err)

		if nil != err {
			logs.Error(
				logs.LOG_TYPE_SYSTEM,
				logs.LOG_KEY_CACHE,
				fmt.Sprintf("failed to get key value. err: %s", err.Error()),
				map[string]interface{}{
					logs.FIELD_KEY_CACHE_KEY: key,
					logs.FIELD_KEY_PAYLOAD:   value,
				},
			)
		}
	}

	return ok
}

func RedisLock(serverPrefix, key string) bool {
	cacheKey := serverPrefix + key
	mutex := getMutex(cacheKey)
	if mutex == nil {
		return false
	}

	err := mutex.Lock()
	if err != nil {
		logs.Error(logs.LOG_TYPE_SYSTEM, logs.LOG_KEY_CACHE, err.Error(),
			map[string]interface{}{
				logs.FIELD_KEY_CACHE_KEY: cacheKey,
			})
		return false
	}

	return true
}

func getMutex(cacheKey string) *redsync.Mutex {
	var mutex *redsync.Mutex
	val, ok := mutexMap.Load(cacheKey)
	if !ok {
		pool := goredis.NewPool(redisConn)
		rs := redsync.New(pool)

		mutex = rs.NewMutex(cacheKey)
		if mutex == nil {
			return nil
		}

		mutexMap.Store(cacheKey, mutex)
	} else {
		mutex, ok = val.(*redsync.Mutex)

		if false == ok {
			logs.Error(
				logs.LOG_TYPE_SYSTEM,
				logs.LOG_KEY_CACHE,
				"Convert redis mutex failed",
				map[string]interface{}{
					logs.FIELD_KEY_CACHE_KEY: cacheKey,
				},
			)
			return nil
		}
	}

	return mutex
}

func RedisUnlock(serverPrefix, key string) bool {
	cacheKey := serverPrefix + key
	mutex := getMutex(cacheKey)
	if nil == mutex {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			"Get redis lock failed",
			map[string]interface{}{
				logs.FIELD_KEY_CACHE_KEY: cacheKey,
			},
		)
		return false
	}

	ok, err := mutex.Unlock()
	if false == ok {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("Unlock redis failed err: %v", err),
			map[string]interface{}{
				logs.FIELD_KEY_CACHE_KEY: key,
			},
		)

		return false
	}

	return true
}
