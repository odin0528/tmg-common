package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
	"xxx/common/configs"
	"xxx/common/logs"
	"xxx/common/utils"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	jsoniter "github.com/json-iterator/go"
)

var (
	redisConn *redis.Client
	mutexMap  sync.Map
	rs        *redsync.Redsync
)

const (
	redisTxPipelineKey string = "redis_tx_pipeline"
	redisPipelineKey   string = "redis_pipeline"
)

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
		return errors.New("new redis client failed")
	}

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		// panic(err)
		return err
	}

	redisConn = redisClient
	pool := goredis.NewPool(redisConn)
	rs = redsync.New(pool)

	return nil
}

func GetRedisClient() *redis.Client {
	if redisConn != nil {
		return redisConn
	}

	return nil
}

func SetRedisClient(client *redis.Client) {
	redisConn = client
}

func ClearApiCenterKey() {
	Delete([]string{RISK_CONTROL_ODDS_TYPE_BACKUP_HASH_KEY})
}

func ClearAll() {
	clearCmsLoginCache()
}

func clearCmsLoginCache() {
	infos, _ := GetAllHashMap(LOGIN_DETAIL_HASH_KEY)

	keys := []string{}
	for key := range infos {
		keys = append(keys, key)
	}

	if len(keys) != 0 {
		DelHashMap(LOGIN_DETAIL_HASH_KEY, keys)
	}

	accountMap, err := GetAllHashMap(LOGIN_ACCOUNT_HASH_KEY)
	if err != nil {
		return
	}

	accountMap["superadmin"] = JWT_CMS_ACCOUNT_TOKEN + "superadmin"

	for account := range accountMap {
		accountCacheKey := JWT_CMS_ACCOUNT_TOKEN + account
		tokenMap, err := GetAllHashMap(accountCacheKey)
		if err != nil {
			continue
		}

		tokens := []string{}
		for token := range tokenMap {
			tokens = append(tokens, token)
		}

		if len(tokens) > 0 {
			if err := DelHashMap(accountCacheKey, tokens); err != nil {
				return
			}
		}
	}
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

func Increase(key string) error {
	return redisConn.Incr(context.Background(), key).Err()
}

func Decrease(key string) (int64, error) {
	return redisConn.Decr(context.Background(), key).Result()
}

func IncreaseBy(key string, value int64) (int64, error) {
	return redisConn.IncrBy(context.Background(), key, value).Result()
}

func IncreaseByFloat(key string, value float64) (float64, error) {
	return redisConn.IncrByFloat(context.Background(), key, value).Result()
}

func Delete(keys []string) error {
	return redisConn.Del(context.Background(), keys...).Err()
}

func GetKey(key string) bool {
	_, ok := redisConn.Get(context.Background(), key).Result()

	if ok == redis.Nil {
		return false
	}

	return true
}

func GetBytes(key string) (retValue []byte, ok bool) {
	cmd := redisConn.Get(context.Background(), key)
	if cmd == nil {
		return nil, false
	}
	value, err := cmd.Bytes()
	if err != nil {
		return nil, false
	}

	return value, true
}

func GetString(key string) (retValue string, ok bool) {
	cmd := redisConn.Get(context.Background(), key)
	if cmd == nil {
		return "", false
	}

	value := cmd.Val()

	return value, true
}

func GetInt(key string) (retValue int, ok bool) {
	cmd := redisConn.Get(context.Background(), key)
	if cmd == nil {
		return 0, false
	}

	value, err := cmd.Int()
	if err != nil {
		return 0, false
	}

	return value, true
}

func GetFloat64(key string) (retValue float64, ok bool) {
	cmd := redisConn.Get(context.Background(), key)
	if cmd == nil {
		return 0, false
	}

	value, err := cmd.Float64()
	if err != nil {
		return 0, false
	}

	return value, true
}

func GetBool(key string) (retValue bool, ok bool) {
	cmd := redisConn.Get(context.Background(), key)
	if cmd == nil {
		return false, false
	}

	byteValues, err := cmd.Bytes()
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
		if value == "" {
			return false
		}

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

func IsExist(key string) bool {
	err := redisConn.Get(context.Background(), key).Err()
	return err == nil
}

func RedisLock(key string, options ...redsync.Option) bool {
	mutex := getMutex(key, options...)
	if mutex == nil {
		return false
	}

	err := mutex.Lock()
	if err != nil {
		logs.Error(logs.LOG_TYPE_SYSTEM, logs.LOG_KEY_CACHE, err.Error(),
			map[string]interface{}{
				logs.FIELD_KEY_CACHE_KEY: key,
			})
		return false
	}

	return true
}

func getMutex(cacheKey string, options ...redsync.Option) *redsync.Mutex {
	if configs.GetBool(configs.SECTION_CACHE, "newmutex", false) {
		var mutex *redsync.Mutex
		val, ok := mutexMap.Load(cacheKey)
		if !ok {
			mutex = rs.NewMutex(cacheKey, options...)
			if mutex == nil {
				return nil
			}

			mutexMap.Store(cacheKey, mutex)
		} else {
			mutex, ok = val.(*redsync.Mutex)

			if !ok {
				logs.Error(
					logs.LOG_TYPE_SYSTEM,
					logs.LOG_KEY_CACHE,
					"convert redis mutex failed",
					map[string]interface{}{
						logs.FIELD_KEY_CACHE_KEY: cacheKey,
					},
				)
				return nil
			}
		}

		return mutex
	} else {
		var mutex *redsync.Mutex
		val, ok := mutexMap.Load(cacheKey)
		if !ok {
			pool := goredis.NewPool(redisConn)
			rs := redsync.New(pool)

			mutex = rs.NewMutex(cacheKey, options...)
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
}

func RedisUnlock(key string) bool {
	mutex := getMutex(key)
	if nil == mutex {
		return false
	}

	ok, _ := mutex.Unlock()

	if false == ok {
		return false
	}

	return true
}

func GetRedisLockOptions(retryDelay time.Duration, retryTimes int, expireSec time.Duration) []redsync.Option {
	return []redsync.Option{
		redsync.WithRetryDelay(retryDelay),
		redsync.WithTries(retryTimes),
		redsync.WithExpiry(expireSec),
	}
}

func BatchRedisLock(redisKeys []string) ([]string, bool) {
	successKeyList := []string{}
	wg := sync.WaitGroup{}
	wg.Add(len(redisKeys))
	for _, key := range redisKeys {
		go func(redisKey string) {
			isLock := RedisLock(redisKey)
			if isLock {
				successKeyList = append(successKeyList, redisKey)
			}

			wg.Done()
		}(key)
	}

	wg.Wait()

	isAllSuccess := false
	if len(successKeyList) == len(redisKeys) {
		isAllSuccess = true
	}

	return successKeyList, isAllSuccess
}

func BatchRedisUnlock(redisKeys []string) ([]string, bool) {
	successKeyList := []string{}
	wg := sync.WaitGroup{}
	wg.Add(len(redisKeys))
	for _, key := range redisKeys {
		go func(redisKey string) {
			isLock := RedisUnlock(redisKey)
			if isLock {
				successKeyList = append(successKeyList, redisKey)
			}

			wg.Done()
		}(key)
	}
	wg.Wait()

	isAllSuccess := false
	if len(successKeyList) == len(redisKeys) {
		isAllSuccess = true
	}

	return successKeyList, isAllSuccess
}

func PutInHashMap(key string, values map[string]interface{}) (err error) {
	setValueMap := map[string]interface{}{}
	for k, v := range values {
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			setValueMap[k] = v
		} else {
			setValueMap[k] = jsonBytes
		}
	}

	return redisConn.HSet(context.Background(), key, setValueMap).Err()
}

// If field not exist, return "redis: nil" error msg
func GetHashMap(key string, field string) (result string, err error) {
	return redisConn.HGet(context.Background(), key, field).Result()
}

func GetAllHashMap(key string) (result map[string]string, err error) {
	return redisConn.HGetAll(context.Background(), key).Result()
}

func DelHashMap(key string, field []string) (err error) {
	return redisConn.HDel(context.Background(), key, field...).Err()
}

func GetHashMapFields(key string) (keys []string, err error) {
	return redisConn.HKeys(context.Background(), key).Result()
}

func IncreaseInHashMap(key, filed string, increaseNum int) (int, error) {
	count, err := redisConn.HIncrBy(context.Background(), key, filed, int64(increaseNum)).Result()
	return int(count), err
}

func LPush(key string, value interface{}) (err error) {
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

	return redisConn.LPush(context.Background(), key, putValue).Err()
}

func RPop(key string) (retValue string, ok bool) {
	cmd := redisConn.RPop(context.Background(), key)
	if cmd == nil {
		return "", false
	}

	value := cmd.Val()

	return value, true
}

func LRange(name string) ([]string, error) {
	return redisConn.LRange(redisConn.Context(), name, 0, -1).Result()
}

func LTrim(name string, start int64, stop int64) {
	redisConn.LTrim(redisConn.Context(), name, start, stop)
}

func Scan(pattern string) ([]string, error) {
	var cursor uint64
	scanAmount := DEFAULT_SCAN_AMOUNT
	result := []string{}

	for {
		var err error
		var keys []string
		keys, cursor, err = redisConn.Scan(redisConn.Context(), cursor, pattern, int64(scanAmount)).Result()
		if err != nil {
			return []string{}, err
		}

		result = append(result, keys...)

		if cursor == 0 {
			break
		}
	}

	return result, nil
}

func HExists(key, field string) (bool, error) {
	return redisConn.HExists(context.Background(), key, field).Result()
}

func HGet(key, field string) (string, error) {
	return redisConn.HGet(context.Background(), key, field).Result()
}

func HSet(key, field, value string) error {
	return redisConn.HSet(context.Background(), key, field, value).Err()
}

func HSetMultiple(key string, data map[string]interface{}) error {
	return redisConn.HSet(context.Background(), key, data).Err()
}

func HGetAll(key string) (map[string]string, error) {
	return redisConn.HGetAll(context.Background(), key).Result()
}

func Keys(key string) ([]string, error) {
	return redisConn.Keys(context.Background(), key).Result()
}

func Expire(key string, timeout time.Duration) (bool, error) {
	return redisConn.Expire(context.Background(), key, timeout).Result()
}

func ExpireNX(key string, timeout time.Duration) (bool, error) {
	return redisConn.ExpireNX(context.Background(), key, timeout).Result()
}

func ExpireXX(key string, timeout time.Duration) (bool, error) {
	return redisConn.ExpireXX(context.Background(), key, timeout).Result()
}

func ExpireGT(key string, timeout time.Duration) (bool, error) {
	return redisConn.ExpireGT(context.Background(), key, timeout).Result()
}

func ExpireLT(key string, timeout time.Duration) (bool, error) {
	return redisConn.ExpireLT(context.Background(), key, timeout).Result()
}

func Publish(channel string, message interface{}) error {
	return redisConn.Publish(context.Background(), channel, message).Err()
}

func Subscribe(channel string) *redis.PubSub {
	return redisConn.Subscribe(context.Background(), channel)
}

type RedisPubSub struct {
	*redis.PubSub
}

func (ps *RedisPubSub) ReceiveMsg() (*redis.Message, error) {
	return ps.ReceiveMessage(context.Background())
}

func IsExistByAllType(key string) bool {
	result, err := redisConn.Exists(context.Background(), key).Result()

	return err == nil && result == 1
}

func SAdd(key string, members ...interface{}) error {
	return redisConn.SAdd(context.Background(), key, members...).Err()
}

func SMembers(key string) ([]string, error) {
	return redisConn.SMembers(context.Background(), key).Result()
}

func SIsMember(key string, member interface{}) (bool, error) {
	return redisConn.SIsMember(context.Background(), key, member).Result()
}

func SPop(key string) (string, error) {
	return redisConn.SPop(context.Background(), key).Result()
}

func SPopN(key string, count int64) ([]string, error) {
	return redisConn.SPopN(context.Background(), key, count).Result()
}

func SRem(key string, members ...interface{}) error {
	return redisConn.SRem(context.Background(), key, members...).Err()
}

func SCard(key string) (int64, error) {
	return redisConn.SCard(context.Background(), key).Result()
}

func SInter(keys ...string) ([]string, error) {
	return redisConn.SInter(context.Background(), keys...).Result()
}

func SUnion(keys ...string) ([]string, error) {
	return redisConn.SUnion(context.Background(), keys...).Result()
}

func SDiff(keys ...string) ([]string, error) {
	return redisConn.SDiff(context.Background(), keys...).Result()
}

func SScanWithCallback(key string, batchSize int64, callback func([]string) error) error {
	var cursor uint64

	for {
		members, nextCursor, err := redisConn.SScan(context.Background(), key, cursor, "", batchSize).Result()
		if err != nil {
			return err
		}

		if err := callback(members); err != nil {
			return err
		}

		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}

	return nil
}

func SPopWithCallback(key string, batchSize int64, callback func([]string) error) error {
	for {
		members, err := redisConn.SPopN(context.Background(), key, batchSize).Result()
		if err != nil {
			return err
		}

		if len(members) == 0 {
			break
		}

		if err := callback(members); err != nil {
			return err
		}
	}

	return nil
}

func SetKeyExpiry(key string, timeout time.Duration) error {
	exists, err := redisConn.Exists(context.Background(), key).Result()
	if err != nil {
		return err
	}

	if exists == 0 {
		return fmt.Errorf("key %s does not exist", key)
	}

	_, err = redisConn.Expire(context.Background(), key, timeout).Result()
	return err
}

func SetKeyExpiryIfNotExist(key string, timeout time.Duration) error {
	success, err := redisConn.ExpireNX(context.Background(), key, timeout).Result()
	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("failed to set expiry: key %s does not exist or already has an expiry time", key)
	}

	return nil
}

func ZIncrBy(key string, increment float64, member string) (float64, error) {
	score, err := redisConn.ZIncrBy(context.Background(), key, increment, member).Result()
	return score, err
}

func ZRangeWithScores(key string, start, stop int64) ([]Z, error) {
	zSlice, err := redisConn.ZRangeWithScores(context.Background(), key, start, stop).Result()
	if err != nil {
		return nil, err
	}

	result := make([]Z, 0, len(zSlice))
	for _, z := range zSlice {
		memberStr, ok := z.Member.(string)
		if ok {
			result = append(result, Z{
				Score:  z.Score,
				Member: memberStr,
			})
		}
	}
	return result, nil
}

func ZRevRangeWithScores(key string, start, stop int64) ([]Z, error) {
	zSlice, err := redisConn.ZRevRangeWithScores(context.Background(), key, start, stop).Result()
	if err != nil {
		return nil, err
	}

	result := make([]Z, 0, len(zSlice))
	for _, z := range zSlice {
		memberStr, ok := z.Member.(string)
		if ok {
			result = append(result, Z{
				Score:  z.Score,
				Member: memberStr,
			})
		}
	}
	return result, nil
}

func ResetPresetRoomGameOption() {
	awaitDeleteKeys := []string{}
	redisPattern := GetCacheKey(GAME_OPTION_PREFIX_KEY, COMMON_LIST_KEY, "*")
	existKeys, err := Scan(redisPattern)
	if err != nil {
		return
	}

	awaitDeleteKeys = append(awaitDeleteKeys, existKeys...)

	batchSize := 1
	for i := 0; i < len(awaitDeleteKeys); i += batchSize {
		end := i + batchSize
		if end > len(awaitDeleteKeys) {
			end = len(awaitDeleteKeys)
		}
		currentDeleteKeys := awaitDeleteKeys[i:end]

		Delete(currentDeleteKeys)
	}
}

func ClearCachedData(pattern string) {
	awaitDeleteKeys := []string{}
	existKeys, err := Scan(pattern)
	if err != nil {
		return
	}

	awaitDeleteKeys = append(awaitDeleteKeys, existKeys...)

	batchSize := 100
	for i := 0; i < len(awaitDeleteKeys); i += batchSize {
		end := i + batchSize
		if end > len(awaitDeleteKeys) {
			end = len(awaitDeleteKeys)
		}
		currentDeleteKeys := awaitDeleteKeys[i:end]

		Delete(currentDeleteKeys)
	}
}

func processPipelineResults(cmds []redis.Cmder) []CmdResult {
	results := make([]CmdResult, 0, len(cmds))
	for _, cmd := range cmds {
		var res CmdResult

		switch c := cmd.(type) {
		case *redis.StringCmd:
			res = CmdResult{
				Result: c.Val(),
				Err:    c.Err(),
			}
		case *redis.IntCmd:
			res = CmdResult{
				Result: strconv.FormatInt(c.Val(), 10),
				Err:    c.Err(),
			}
		case *redis.FloatCmd:
			res = CmdResult{
				Result: strconv.FormatFloat(c.Val(), 'f', -1, 64),
				Err:    c.Err(),
			}
		case *redis.BoolCmd:
			res = CmdResult{
				Result: strconv.FormatBool(c.Val()),
				Err:    c.Err(),
			}
		case *redis.StringSliceCmd:
			val, err := c.Result()
			if err != nil {
				res = CmdResult{
					Result: "",
					Err:    err,
				}
			} else {
				jsonBytes, err := json.Marshal(val)
				if err != nil {
					res = CmdResult{
						Result: "",
						Err:    err,
					}
				} else {
					res = CmdResult{
						Result: string(jsonBytes),
						Err:    nil,
					}
				}
			}
		case *redis.IntSliceCmd:
			val, err := c.Result()
			if err != nil {
				res = CmdResult{
					Result: "",
					Err:    err,
				}
			} else {
				jsonBytes, err := json.Marshal(val)
				if err != nil {
					res = CmdResult{
						Result: "",
						Err:    err,
					}
				} else {
					res = CmdResult{
						Result: string(jsonBytes),
						Err:    nil,
					}
				}
			}
		case *redis.StatusCmd:
			res = CmdResult{
				Result: c.Val(),
				Err:    c.Err(),
			}
		case *redis.SliceCmd:
			val, err := c.Result()
			if err != nil {
				res = CmdResult{
					Result: "",
					Err:    err,
				}
			} else {
				jsonBytes, err := json.Marshal(val)
				if err != nil {
					res = CmdResult{
						Result: "",
						Err:    err,
					}
				} else {
					res = CmdResult{
						Result: string(jsonBytes),
						Err:    nil,
					}
				}
			}
		case *redis.StringStringMapCmd:
			val, err := c.Result()
			if err != nil {
				res = CmdResult{
					Result: "",
					Err:    err,
				}
			} else {
				jsonBytes, err := json.Marshal(val)
				if err != nil {
					res = CmdResult{
						Result: "",
						Err:    err,
					}
				} else {
					res = CmdResult{
						Result: string(jsonBytes),
						Err:    nil,
					}
				}
			}
		default:
			res = CmdResult{
				Result: cmd.String(),
				Err:    cmd.Err(),
			}
		}

		results = append(results, res)
	}
	return results
}

// RunPipelined 使用回調函數執行 Pipeline 操作
func RunPipelined(ctx context.Context, fn func(Pipeline)) ([]CmdResult, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, errors.New("redis client not initialized")
	}

	pipe := client.Pipeline()

	wrapper := &PipelineWrapper{
		pipe: pipe,
	}

	// 使用者自定義要執行哪些 pipeline 操作
	fn(wrapper)

	cmds, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	return processPipelineResults(cmds), nil
}

func GetPipeline() Pipeline {
	client := GetRedisClient()
	pipe := client.Pipeline()

	wrapper := &PipelineWrapper{
		pipe: pipe,
	}

	return wrapper
}

func WithPipeline(ctx context.Context, pipeline Pipeline) context.Context {
	return context.WithValue(ctx, redisPipelineKey, pipeline)
}

func GetPipelineWithContext(ctx context.Context) (Pipeline, bool) {
	pipeline, ok := ctx.Value(redisPipelineKey).(Pipeline)
	return pipeline, ok
}

func GetTxPipeline() TxPipeline {
	client := GetRedisClient()
	pipe := client.TxPipeline()

	wrapper := &TxPipelineWrapper{
		PipelineWrapper{pipe},
	}

	return wrapper
}

func WithTxPipeline(ctx context.Context, pipeline TxPipeline) context.Context {
	return context.WithValue(ctx, redisTxPipelineKey, pipeline)
}

func GetTxPipelineWithContext(ctx context.Context) (TxPipeline, bool) {
	txPipeline, ok := ctx.Value(redisTxPipelineKey).(TxPipeline)
	return txPipeline, ok
}

//func RunTxPipelineWithCtx(ctx context.Context) ([]CmdResult, error) {
//	pipeTx, ok := GetTxPipelineWithContext(ctx)
//	if !ok {
//		return nil, errors.New("no tx pipeline found")
//	}
//
//	cmds, err := pipeTx.Exec(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	return processPipelineResults(cmds), nil
//}

func DiscardTxPipeline(ctx context.Context) error {
	pipeTx, ok := GetTxPipelineWithContext(ctx)
	if !ok {
		return errors.New("no tx pipeline found")
	}

	return pipeTx.Discard()
}

func AddMultiZSetByScriptBuckets(buckets map[string][]any) error {
	_, err := RunPipelined(context.Background(), func(p Pipeline) {
		for key, raws := range buckets {
			if len(raws) == 0 {
				continue
			}

			var zs []*redis.Z
			for _, raw := range raws {
				b, err := json.Marshal(raw)
				if err != nil {
					continue
				}
				var m map[string]any
				if err := json.Unmarshal(b, &m); err != nil {
					continue
				}
				idVal, ok := m["id"].(float64)
				if !ok {
					continue
				}
				zs = append(zs, &redis.Z{
					Score:  idVal,
					Member: b,
				})
			}
			p.ZAdd(key, zs...)
		}
	})
	return err
}

func ZCard(key string) (int64, error) {
	count, err := redisConn.ZCard(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func ZRange(key string, start, stop int64) ([]string, error) {
	result, err := redisConn.ZRange(context.Background(), key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteWithContext(ctx context.Context, keys []string) error {
	const batchSize = 5000
	for i := 0; i < len(keys); i += batchSize {
		end := i + batchSize
		if end > len(keys) {
			end = len(keys)
		}

		batch := keys[i:end]
		if err := redisConn.Del(ctx, batch...).Err(); err != nil {
			return err
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
	return nil
}

func AccessLimit(ctx context.Context, key string, interval int, limit int) (cnt int64, err error) {
	accessLimitScript := redis.NewScript(`
	local count = redis.call('INCR', KEYS[1])
	if count == 1 then
	redis.call('EXPIRE', KEYS[1], ARGV[1])
	end
	return count`)

	cnt, err = accessLimitScript.Run(ctx, redisConn, []string{key}, interval).Int64()
	if err != nil {
		return 0, err
	}

	if cnt >= int64(limit) {
		return cnt, fmt.Errorf("access limit exceeded: %d", cnt)
	}

	return cnt, nil
}

func NewLuaScript(script string) *redis.Script {
	return redis.NewScript(script)
}

func RunScript(script *redis.Script, keys []string, args ...interface{}) (interface{}, error) {
	return script.Run(context.Background(), redisConn, keys, args...).Result()
}

func XAdd(key string, values map[string]interface{}, maxLenApprox int64) (string, error) {
	return redisConn.XAdd(context.Background(), &redis.XAddArgs{
		Stream:       key,
		ID:           "*",
		Values:       values,
		MaxLenApprox: maxLenApprox,
	}).Result()
}

func XRange(key string, startTime, endTime int64, limit int64) ([]redis.XMessage, error) {
	start := fmt.Sprintf("%d-0", startTime)
	end := fmt.Sprintf("%d-0", endTime)
	if limit == 0 {
		return redisConn.XRange(context.Background(), key, start, end).Result()
	}
	return redisConn.XRangeN(context.Background(), key, start, end, limit).Result()
}

func XRevRange(key string, startTime, endTime int64, limit int64) ([]redis.XMessage, error) {
	start := fmt.Sprintf("%d-0", startTime)
	end := fmt.Sprintf("%d-0", endTime)
	if limit == 0 {
		return redisConn.XRevRange(context.Background(), key, start, end).Result()
	}
	return redisConn.XRevRangeN(context.Background(), key, start, end, limit).Result()
}

func XLen(key string) (int64, error) {
	return redisConn.XLen(context.Background(), key).Result()
}

func XTrimResult(key string, expireTime time.Duration, trimLimit int64) (int64, error) {
	cutoffID := fmt.Sprintf("%d-0", time.Now().
		Add(-expireTime).
		UnixMilli(),
	)

	return redisConn.XTrimMinIDApprox(context.Background(), key, cutoffID, trimLimit).Result()
}

func XTrim(key string, expireTime time.Duration, trimLimit int64) error {
	cutoffID := fmt.Sprintf("%d-0", time.Now().
		Add(-expireTime).
		UnixMilli(),
	)

	return redisConn.XTrimMinIDApprox(context.Background(), key, cutoffID, trimLimit).Err()
}

func GetKeyValue(pattern string) (map[string]interface{}, error) {
	cursor := uint64(0)
	result := make(map[string]interface{})

	ctx := redisConn.Context()

	for {
		keys, nextCursor, err := redisConn.Scan(ctx, cursor, pattern, DEFAULT_SCAN_AMOUNT).Result()
		if err != nil {
			return result, err
		}

		for _, key := range keys {
			keyType, err := redisConn.Type(ctx, key).Result()
			if err != nil {
				continue
			}

			switch keyType {
			case "string":
				val, err := redisConn.Get(ctx, key).Result()
				if err == nil {
					result[key] = val
				}
			case "hash":
				val, err := redisConn.HGetAll(ctx, key).Result()
				if err == nil {
					result[key] = val
				}
			case "list":
				val, err := redisConn.LRange(ctx, key, 0, -1).Result()
				if err == nil {
					result[key] = val
				}
			case "set":
				val, err := redisConn.SMembers(ctx, key).Result()
				if err == nil {
					result[key] = val
				}
			case "zset":
				val, err := redisConn.ZRangeWithScores(ctx, key, 0, -1).Result()
				if err == nil {
					result[key] = val
				}
			default:
				result[key] = fmt.Sprintf("[Unsupported type: %s]", keyType)
			}
		}

		if nextCursor == 0 {
			break
		}
		cursor = nextCursor
	}

	return result, nil
}

func GetValue(key string) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	ctx := redisConn.Context()

	keyType, err := redisConn.Type(ctx, key).Result()
	if err != nil {
		return result, err
	}

	switch keyType {
	case "string":
		val, err := redisConn.Get(ctx, key).Result()
		if err == nil {
			result[key] = val
		}
	case "hash":
		val, err := redisConn.HGetAll(ctx, key).Result()
		if err == nil {
			result[key] = val
		}
	case "list":
		val, err := redisConn.LRange(ctx, key, 0, -1).Result()
		if err == nil {
			result[key] = val
		}
	case "set":
		val, err := redisConn.SMembers(ctx, key).Result()
		if err == nil {
			result[key] = val
		}
	case "zset":
		val, err := redisConn.ZRangeWithScores(ctx, key, 0, -1).Result()
		if err == nil {
			result[key] = val
		}
	default:
		result[key] = fmt.Sprintf("[Unsupported type: %s]", keyType)
	}
	return result, nil
}

func DeleteKeys(pattern string) error {
	var (
		cursor  uint64
		deleted int64
		err     error
	)

	key := pattern
	ctx := redisConn.Context()

	if strings.ContainsAny(pattern, "*?[") {
		for {
			var keys []string
			keys, cursor, err = redisConn.Scan(ctx, cursor, key, DEFAULT_SCAN_AMOUNT).Result()
			if err != nil {
				return err
			}
			if len(keys) > 0 {
				pipe := redisConn.Pipeline()
				for _, k := range keys {
					pipe.Del(ctx, k)
				}
				_, err := pipe.Exec(ctx)
				if err != nil {
					return err
				}
				deleted += int64(len(keys))
			}
			if cursor == 0 {
				break
			}
		}
	} else {
		// 精確刪除
		_, err = redisConn.Del(ctx, key).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

func ClearAgentIdSerialNum() {
	now := time.Now()

	reserveListMap := map[string]bool{}

	for i := 0; i < CLEAR_AGENT_ID_SERIAL_NUM_X_DAYS_BEFORE; i++ {
		reserveListMap[fmt.Sprintf("%s%s", AGENT_ID_SERIAL_NUMBER_PREFIX, now.AddDate(0, 0, -i).Format(utils.DATE_YYYYMMDD_FORMAT))] = true
	}

	currntKeyList, err := Scan(AGENT_ID_SERIAL_NUMBER_PREFIX + "*")
	if err != nil {
		return
	}

	deleteKeys := []string{}
	for _, key := range currntKeyList {
		if !reserveListMap[key] {
			deleteKeys = append(deleteKeys, key)
		}
	}

	ctx := context.Background()
	if len(deleteKeys) > 0 {
		pipe := redisConn.Pipeline()
		for _, k := range deleteKeys {
			pipe.Del(ctx, k)
		}

		pipe.Exec(ctx)
	}
}

func SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return redisConn.SetNX(context.Background(), key, value, expiration).Result()
}
