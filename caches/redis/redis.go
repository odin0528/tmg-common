// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package redis for cache provider
//
// depend on github.com/gomodule/redigo/redis
//
// go install github.com/gomodule/redigo/redis
//
// Usage:
// import(
//
//	_ "github.com/astaxie/beego/cache/redis"
//	"github.com/astaxie/beego/cache"
//
// )
//
//	bm, err := cache.NewCache("redis", `{"conn":"127.0.0.1:11211"}`)
//
//	more docs http://beego.me/docs/module/cache.md
package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"sync"
	"time"
	"xxx/common/caches"
	"xxx/common/configs"
	"xxx/common/logs"

	"github.com/alicebob/miniredis"
	"github.com/gomodule/redigo/redis"

	"strings"
)

var (
	// DefaultKey the collection name of redis for cache adapter.
	DefaultKey = "beecacheRedis"
)

// Cache is Redis cache adapter.
type Cache struct {
	p        *redis.Pool // redis connection pool
	conninfo string
	dbNum    int
	key      string
	password string
	maxIdle  int
	useTLS   bool
}

type reply struct {
	err   error
	value interface{}
}

// NewRedisCache create new redis cache with default collection name.
func NewRedisCache() caches.Cache {
	return &Cache{key: DefaultKey}
}

func (rc *Cache) GetPool() *redis.Pool {
	return rc.p
}

// actually do the redis cmds, args[0] must be the key name.
func (rc *Cache) do(commandName string, args ...interface{}) (reply interface{}, err error) {
	if len(args) < 1 {
		return nil, errors.New("missing required arguments")
	}
	args[0] = rc.associate(args[0])
	c := rc.p.Get()

	defer c.Close()

	return c.Do(commandName, args...)
}

func (rc *Cache) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return rc.do(commandName, args...)
}

// associate with config key.
func (rc *Cache) associate(originKey interface{}) string {
	return fmt.Sprintf("%s:%s", rc.key, originKey)
}

var iterMutex sync.Mutex

func (rc *Cache) Keys(pattern string, count int) []string {
	c := rc.p.Get()
	defer func() {
		iterMutex.Unlock()
		c.Close()
	}()

	iterMutex.Lock()
	var iter = 0
	redisPrefix := configs.Get("cache", "prefix_key", "zio_center_redis") + ":"
	keys := []string{}
	retKeys := []string{}
	keyNum := 0
	scanCount := count

	if 0 == count {
		scanCount = 1000
	}

	for {
		if arr, err := redis.MultiBulk(c.Do("SCAN", iter, "MATCH", pattern, "COUNT", scanCount)); err != nil {
			logs.Error(logs.LOG_TYPE_SYSTEM, logs.LOG_KEY_CACHE, "Keys MultiBulk error", map[string]interface{}{
				logs.FIELD_KEY_PAYLOAD: err.Error(),
			})
			return []string{}
		} else {
			iter, _ = redis.Int(arr[0], nil)
			keys, _ = redis.Strings(arr[1], nil)
		}

		for _, key := range keys {
			retKeys = append(retKeys, strings.TrimPrefix(key, redisPrefix))
			keyNum++

			if 0 != count && count <= keyNum {
				break
			}
		}

		if 0 == iter || (0 != count && count <= keyNum) {
			break
		}
	}

	return retKeys
}

// Get cache from redis.
func (rc *Cache) Get(key string) interface{} {
	if v, err := rc.do("GET", key); err == nil {
		return v
	}
	return nil
}

// GetMulti get cache from redis.
func (rc *Cache) GetMulti(keys []string) []interface{} {
	c := rc.p.Get()
	defer c.Close()
	var args []interface{}
	for _, key := range keys {
		args = append(args, rc.associate(key))
	}
	values, err := redis.Values(c.Do("MGET", args...))
	if err != nil {
		return nil
	}
	return values
}

// Put put cache to redis.
func (rc *Cache) Put(key string, val interface{}, timeout time.Duration) error {
	_, err := rc.do("SETEX", key, int64(timeout/time.Second), val)
	if nil != err {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to put key. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
				logs.FIELD_KEY_PAYLOAD:    val,
			},
		)
	}
	return err
}

// PutNoExpiry put cache to redis without timeout.
func (rc *Cache) PutNoExpiry(key string, val interface{}) error {
	_, err := rc.do("SET", key, val)
	if nil != err {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to put key with no expiry. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
				logs.FIELD_KEY_PAYLOAD:    val,
			},
		)
	}
	return err
}

// Delete delete cache in redis.
func (rc *Cache) Delete(key string) error {
	_, err := rc.do("DEL", key)
	if nil != err {
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

// LPush push value into tail of LIST
func (rc *Cache) LPush(key string, val interface{}) error {
	_, err := rc.do("LPUSH", key, val)
	return err
}

// RPop return value head of LIST
func (rc *Cache) RPop(key string) interface{} {
	v, err := rc.do("RPOP", key)
	if err == nil {
		return v
	}
	return nil
}

// LLen return value of len of LIST
func (rc *Cache) LLen(key string) *int {
	if val, err := rc.do("LLEN", key); err == nil {
		if v, err := redis.Int(val, err); nil == err {
			return &v
		}
		return nil
	}
	return nil
}

func (rc *Cache) LRange(key string, start, end int) []interface{} {
	values, err := redis.Values(rc.Do("LRANGE", key, start, end))
	if err != nil {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to LRange key err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
				logs.FIELD_KEY_START:      start,
				logs.FIELD_KEY_END:        end,
			},
		)
		return nil
	}
	return values
}

// LTrim it will contain only the specified range of elements specified
func (rc *Cache) LTrim(key string, start, end int) error {
	_, err := rc.do("LTRIM", key, start, end)
	return err
}

func (rc *Cache) HMSet(key string, values map[string]interface{}) error {
	for k, v := range values {
		json, _ := json.Marshal(v)
		values[k] = json
	}
	args := redis.Args{key}.AddFlat(values)
	_, err := rc.Do("HMSET", args...)
	return err
}

func (rc *Cache) HMGet(key string, fields []string) []string {
	//_ = rc.HGetAll(key)
	values, err := redis.Strings(rc.Do("HMGet", redis.Args{key}.AddFlat(fields)...))
	if err != nil {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to HMGet. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
				logs.FIELD_KEY_PAYLOAD:    fields,
			},
		)

		return nil
	}

	return values
}

func (rc *Cache) HGetAll(key string) map[string]string {
	values, err := redis.StringMap(rc.Do("HGETALL", key))
	if err != nil {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to HGetAll. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
			},
		)

		return nil
	}
	return values
}

func (rc *Cache) HDEL(key string, fields []string) bool {
	_, err := rc.Do("HDEL", redis.Args{key}.AddFlat(fields)...)
	if err != nil {
		logs.Error(
			logs.LOG_TYPE_SYSTEM,
			logs.LOG_KEY_CACHE,
			fmt.Sprintf("failed to HDEL. err: %s", err.Error()),
			map[string]interface{}{
				logs.FIELD_KEY_REDIST_KEY: key,
				logs.FIELD_KEY_PAYLOAD:    fields,
			},
		)

		return false
	}

	return true
}

func (rc *Cache) DeleteMulti(keys []string) error {
	if 0 == len(keys) {
		return nil
	}
	var args []interface{}
	for _, key := range keys {
		args = append(args, rc.associate(key))
	}

	c := rc.p.Get()

	defer c.Close()

	_, err := c.Do("DEL", args...)
	return err
}

func (rc *Cache) DeleteMultiWithoutPrefix(keys []string) error {
	var args []interface{}
	for _, key := range keys {
		args = append(args, key)
	}

	c := rc.p.Get()

	defer c.Close()

	_, err := c.Do("DEL", args...)

	return err
}

// IsExist check cache's existence in redis.
func (rc *Cache) IsExist(key string) bool {
	v, err := redis.Bool(rc.do("EXISTS", key))
	if err != nil {
		return false
	}
	return v
}

// Incr increase counter in redis.
func (rc *Cache) Incr(key string) error {
	_, err := redis.Bool(rc.do("INCRBY", key, 1))
	return err
}

// Decr decrease counter in redis.
func (rc *Cache) Decr(key string) error {
	_, err := redis.Bool(rc.do("INCRBY", key, -1))
	return err
}

// Incr increase float in redis.
func (rc *Cache) IncrFloat(key string, value float64) error {
	_, err := redis.Bool(rc.do("INCRBYFLOAT", key, value))
	return err
}

// ClearAll clean all cache in redis. delete this redis collection.
func (rc *Cache) ClearAll() error {
	c := rc.p.Get()
	defer c.Close()
	cachedKeys, err := redis.Strings(c.Do("KEYS", rc.key+":*"))
	if err != nil {
		return err
	}
	for _, str := range cachedKeys {
		if _, err = c.Do("DEL", str); err != nil {
			return err
		}
	}
	return err
}

// StartAndGC start redis cache adapter.
// config is like {"key":"collection key","conn":"connection info","dbNum":"0"}
// the cache item in redis are stored forever,
// so no gc operation.
func (rc *Cache) StartAndGC(config string) error {
	var cf map[string]string
	json.Unmarshal([]byte(config), &cf)

	if _, ok := cf["key"]; !ok {
		cf["key"] = DefaultKey
	}
	if _, ok := cf["conn"]; !ok {
		return errors.New("config has no conn key")
	}

	// Format redis://<password>@<host>:<port>
	cf["conn"] = strings.Replace(cf["conn"], "redis://", "", 1)
	if i := strings.Index(cf["conn"], "@"); i > -1 {
		cf["password"] = cf["conn"][0:i]
		cf["conn"] = cf["conn"][i+1:]
	}

	if _, ok := cf["dbNum"]; !ok {
		cf["dbNum"] = "0"
	}
	if _, ok := cf["password"]; !ok {
		cf["password"] = ""
	}
	if _, ok := cf["maxIdle"]; !ok {
		cf["maxIdle"] = "100"
	}
	if _, ok := cf["use_tls"]; !ok {
		cf["use_tls"] = "no"
	}

	rc.key = cf["key"]
	rc.conninfo = cf["conn"]
	rc.dbNum, _ = strconv.Atoi(cf["dbNum"])
	rc.password = cf["password"]
	rc.maxIdle, _ = strconv.Atoi(cf["maxIdle"])

	if "yes" == cf["use_tls"] {
		rc.useTLS = true
	}

	rc.connectInit()

	c := rc.p.Get()
	defer c.Close()

	return c.Err()
}

// connect to redis.
func (rc *Cache) connectInit() {
	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", rc.conninfo, redis.DialUseTLS(rc.useTLS))
		if err != nil {
			return nil, err
		}

		if rc.password != "" {
			if _, err := c.Do("AUTH", rc.password); err != nil {
				c.Close()
				return nil, err
			}
		}

		_, selecterr := c.Do("SELECT", rc.dbNum)
		if selecterr != nil {
			c.Close()
			return nil, selecterr
		}
		return
	}
	// initialize a new pool
	rc.p = &redis.Pool{
		MaxIdle:     rc.maxIdle,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}
}

func init() {
	caches.Register("redis", NewRedisCache)
}

func (rc *Cache) Mock(s *miniredis.Miniredis) {
	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", s.Addr())
		if err != nil {
			return nil, err
		}

		return
	}
	rc.p = &redis.Pool{
		MaxIdle:     rc.maxIdle,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}
}

func (rc *Cache) HEXISTS(key string, field string) (bool, error) {
	if v, err := rc.do("HEXISTS", key, field); err == nil {
		value := reflect.ValueOf(v).Int()
		if value == 1 {
			return true, nil
		}
		return false, nil
	} else {
		return false, err
	}
}

func (rc *Cache) HSET(key string, field string, value interface{}) error {
	_, err := rc.do("HSET", key, field, value)

	return err
}

func (rc *Cache) HGET(key string, field string) (interface{}, error) {
	v, err := rc.do("HGET", key, field)

	if nil == v {
		return v, errors.New("Cache Not Exsit")
	}

	return v, err
}

func (rc *Cache) Pipeline(commandSlice [][]string) (reply interface{}, err error) {
	if len(commandSlice) < 1 {
		return nil, errors.New("Pipeline missing required arguments")
	}

	conn := rc.p.Get()

	defer conn.Close()

	conn.Send("MULTI")

	for idx, args := range commandSlice {
		argsLen := len(args)

		if argsLen < 2 {
			return nil, errors.New("error args idx: " + fmt.Sprintf("idx: %d\nargs: %s\n", idx, args))
		}

		command := args[0]           // redis command
		key := rc.associate(args[1]) // binding redis prefixKey

		conn.Send(command, redis.Args{key}.AddFlat(args[2:])...)
	}

	return conn.Do("EXEC")
}
