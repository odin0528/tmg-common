package cache

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"
	cache "xxx/common/caches"
	"xxx/common/caches/redis"
	"xxx/common/configs"
	"xxx/common/logs"

	"github.com/alicebob/miniredis"
	"github.com/go-redsync/redsync"
	jsoniter "github.com/json-iterator/go"
)

var cacheAdapter cache.Cache = nil
var mutexMap sync.Map

func getRedisSettings() string {
	section := "cache"
	setting := fmt.Sprintf("{\"key\":\"%s\", \"conn\":\"%s:%s\", \"dbNum\":\"%s\", \"password\":\"%s\", \"use_tls\":\"%s\"}",
		configs.Get(section, "prefix_key", ""),
		configs.Get(section, "host", "localhost"),
		configs.Get(section, "port", "6379"),
		configs.Get(section, "dbNum", "0"),
		configs.Get(section, "password", ""),
		configs.Get(section, "use_tls", "no"),
	)

	return setting
}

// Cache will be created with config/cache.ini file.
// Please refer https://beego.me/docs/module/cache.md to fill the config.
// Support for memory, file, redis and memcache as cache engine.
func ConfigNewCache() {
	var err error
	cacheAdapter, err = cache.NewCache(configs.Get("cache", "engine", ""), getRedisSettings())

	if err != nil {
		panic(err)
	}
}

func MockRedis() {
	s, err := miniredis.Run()
	if nil != err {
		panic(err)
	}
	cacheAdapter, err = cache.NewMockCache(s)
	if nil != err {
		panic(err)
	}
}

func LPush(key string, val interface{}) error {
	err := cacheAdapter.LPush(key, val)
	if err != nil {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to LPush key. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
				logs.FIELD_KEY_PAYLOAD:    val,
			},
		)
	}

	return err
}

func RPop(key string) interface{} {
	return cacheAdapter.RPop(key)
}

func LLen(key string) (int, error) {
	val := cacheAdapter.LLen(key)
	if nil == val {
		err := errors.New("cacheAdapter.LLen failed")
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to LLen key. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
			},
		)
		return -1, err
	}

	return *val, nil
}

func LRange(key string, start, end int) []interface{} {
	return cacheAdapter.LRange(key, start, end)
}

func LTrim(key string, start, end int) error {
	err := cacheAdapter.LTrim(key, start, end)
	if err != nil {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to LTrim key. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
				logs.FIELD_KEY_START:      start,
				logs.FIELD_KEY_END:        end,
			},
		)
	}

	return err
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
					logs.FIELD_KEY_REDIST_KEY: key,
					logs.FIELD_KEY_PAYLOAD:    value,
					logs.FIELD_KEY_TIMEOUT:    timeout,
				},
			)

			return err
		}
	}

	err = cacheAdapter.Put(key, putValue, timeout)
	if err != nil {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to put key. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
				logs.FIELD_KEY_PAYLOAD:    value,
				logs.FIELD_KEY_TIMEOUT:    timeout,
			},
		)
	}

	return err
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
					logs.FIELD_KEY_REDIST_KEY: key,
					logs.FIELD_KEY_PAYLOAD:    value,
				},
			)
			return err
		}
	}

	err = cacheAdapter.PutNoExpiry(key, putValue)
	if err != nil {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to put key. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
				logs.FIELD_KEY_PAYLOAD:    value,
			},
		)
	}

	return err
}

func Delete(key string) error {
	err := cacheAdapter.Delete(key)
	if err != nil {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to delete key. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
			},
		)

	}

	return err
}

func ClearAll() error {
	err := cacheAdapter.ClearAll()
	if err != nil {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to clear all key. err: %s", err.Error()),
			map[string]interface{}{},
		)
	}

	return err
}

func Get(key string) interface{} {
	return cacheAdapter.Get(key)
}

func GetString(key string) (retValue string, ok bool) {
	if value := cacheAdapter.Get(key); nil != value {
		retValue = cache.GetString(value)
		ok = true
	}

	return retValue, ok
}

func GetBool(key string) (retValue bool, ok bool) {
	if value := cacheAdapter.Get(key); nil != value {
		retValue = cache.GetBool(value)
		ok = true
	}

	return retValue, ok
}

func GetFloat64(key string) (retValue float64, ok bool) {
	if value := cacheAdapter.Get(key); nil != value {
		retValue = cache.GetFloat64(value)
		ok = true
	}

	return retValue, ok
}

func GetInt(key string) (retValue int, ok bool) {
	if value := cacheAdapter.Get(key); nil != value {
		retValue = cache.GetInt(value)
		ok = true
	}

	return retValue, ok
}

func GetInt64(key string) (retValue int64, ok bool) {
	if value := cacheAdapter.Get(key); nil != value {
		retValue = cache.GetInt64(value)
		ok = true
	}

	return retValue, ok
}

func GetStructData(key string, data interface{}) bool {
	value, ok := GetString(key)

	if true == ok {
		json := jsoniter.ConfigCompatibleWithStandardLibrary
		err := json.Unmarshal(([]byte)(value), data)
		ok = (nil == err)

		if nil != err {
			logs.Error(
				logs.LOG_TYPE_SYSTEM,
				logs.LOG_KEY_CACHE,
				fmt.Sprintf("failed to get key value. err: %s", err.Error()),
				map[string]interface{}{
					logs.FIELD_KEY_REDIST_KEY: key,
					logs.FIELD_KEY_PAYLOAD:    value,
				},
			)
		}
	}

	return ok
}

func Increase(key string) bool {
	err := cacheAdapter.Incr(key)
	if nil != err {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to increase key. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
			},
		)

		return false
	}

	return true
}

func Decrease(key string) bool {
	err := cacheAdapter.Decr(key)
	if nil != err {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to decrease key. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
			},
		)
		return false
	}

	return true
}

func IncreaseFloat(key string, value float64) bool {
	err := cacheAdapter.IncrFloat(key, value)
	if nil != err {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to increase float key. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
				logs.FIELD_KEY_PAYLOAD:    value,
			},
		)
		return false
	}

	return true
}

// Put count as 0 will load all keys from the pattern
func Keys(pattern string, count int) []string {
	return cacheAdapter.Keys(pattern, count)
}

func DeleteMulti(keys []string) error {
	err := cacheAdapter.DeleteMulti(keys)
	if nil != err {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to delete multiple key. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: keys,
			},
		)
	}

	return err
}

func DeleteMultiWithoutPrefix(keys []string) error {
	err := cacheAdapter.DeleteMultiWithoutPrefix(keys)
	if nil != err {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to delete multiple with prefix key. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: keys,
			},
		)
	}

	return err
}

func IsExist(key string) bool {
	return cacheAdapter.IsExist(key)
}

func GetRedisCache() (*redis.Cache, bool) {
	cache, ok := cacheAdapter.(*redis.Cache)

	if false == ok {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			"failed to get redis handler.",
			map[string]interface{}{},
		)

	}

	return cache, ok
}

func RedisLock(key string) bool {
	mutex := getMutex(key)

	if nil == mutex {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			"Get redis lock failed",
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
			},
		)

		return false
	}

	err := mutex.Lock()

	if nil != err {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			"Lock redis failed",
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
			},
		)
		return false
	}

	return true
}

func BatchRedisLock(keys []string) bool {
	var wait sync.WaitGroup
	var isAllSuccess bool = true
	wait.Add(len(keys))

	lockFunc := func(key string, wait *sync.WaitGroup) {
		ok := RedisLock(key)

		if false == ok {
			logs.Error(
				logs.LOG_TYPE_SYSTEM,
				logs.LOG_KEY_CACHE,
				"Batch redis lock failed",
				map[string]interface{}{
					logs.FIELD_KEY_REDIST_KEY: key,
				},
			)
			isAllSuccess = false
		}

		(*wait).Done()
	}

	for _, key := range keys {
		go lockFunc(key, &wait)
	}

	wait.Wait()

	return isAllSuccess
}

func RedisUnlock(key string) bool {
	mutex := getMutex(key)

	if nil == mutex {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			"Get redis lock failed",
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
			},
		)
		return false
	}

	ok, err := mutex.Unlock()

	if false == ok {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("Unlock redis failed err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
			},
		)

		return false
	}

	return true
}

func BatchRedisUnlock(keys []string) bool {
	if 0 == len(keys) {
		return true
	}

	unlockKeys := []string{}

	for _, key := range keys {
		unlockKeys = append(unlockKeys, WOW_MUTEX_PREFIX+key)
	}

	err := DeleteMultiWithoutPrefix(unlockKeys)

	if nil != err {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("Batch unlock redis failed err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: keys,
			},
		)
	}

	return nil == err
}

func getMutex(key string) *redsync.Mutex {
	cacheKey := WOW_MUTEX_PREFIX + key
	var mutex *redsync.Mutex
	val, ok := mutexMap.Load(cacheKey)

	if false == ok {
		distributedLock := getDistributedLock()

		if nil == distributedLock {
			return nil
		}

		mutex = distributedLock.NewMutex(cacheKey, redsync.SetExpiry(time.Second*MUTEX_DURATION_SECOND))

		if nil == mutex {
			logs.Error(
				logs.LOG_TYPE_SYSTEM,
				logs.LOG_KEY_CACHE,
				"distributedLock.NewMutex failed",
				map[string]interface{}{
					logs.FIELD_KEY_REDIST_KEY: key,
				},
			)

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
					logs.FIELD_KEY_REDIST_KEY: key,
				},
			)
			return nil
		}
	}

	return mutex
}

func getDistributedLock() *redsync.Redsync {
	redis, ok := GetRedisCache()

	if false == ok {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			"getDistributedLock failed",
			map[string]interface{}{},
		)

		return nil
	}

	return redsync.New([]redsync.Pool{redis.GetPool()})
}

func PutInHashMap(key string, values map[string]interface{}) bool {
	if err := cacheAdapter.HMSet(key, values); nil != err {
		if nil != err {
			logs.Error(
				logs.LOG_TYPE_SYSTEM,
				logs.LOG_KEY_CACHE,
				fmt.Sprintf("Failed to PutInHashMap err:%s", err.Error()),
				map[string]interface{}{
					logs.FIELD_KEY_REDIST_KEY: key,
					logs.FIELD_KEY_PAYLOAD:    values,
				},
			)

			return false
		}
	}

	return true
}

func DelHashMap(key string, fields []string) bool {
	return cacheAdapter.HDEL(key, fields)
}

func GetHashMap(key string) map[string]string {
	return cacheAdapter.HGetAll(key)
}

func GetHashMapWithFields(key string, fields []string) []string {
	return cacheAdapter.HMGet(key, fields)
}

func HashMapFieldExsit(key string, field string) (bool, error) {
	return cacheAdapter.HEXISTS(key, field)
}

func HashMapSetField(key string, field string, value interface{}) error {
	return cacheAdapter.HSET(key, field, value)
}

func HashMapGetField(key string, field string) ([]byte, bool) {
	val, err := cacheAdapter.HGET(key, field)

	if nil != err {
		return []byte{}, false
	}

	return reflect.ValueOf(val).Bytes(), true
}

// 會占用 pipeline memory
// example
// [
//
//	["SET", "KEY", "VALUE"],
//	["GET", "KEY"],
//
// ]
func Pipeline(commandSlice [][]string) (res [][]byte, err error) {
	replys, err := cacheAdapter.Pipeline(commandSlice)

	if nil != err {
		return [][]byte{}, err
	}

	r := [][]byte{}
	for _, reply := range replys.([]interface{}) {
		if bytes, ok := reply.([]byte); ok {
			r = append(r, bytes)
		}
	}

	return r, nil
}
