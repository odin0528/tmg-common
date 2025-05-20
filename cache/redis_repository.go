package cache

import (
	"mgmt/common/redis"
	"time"

	"github.com/go-redsync/redsync/v4"
)

func NewRedisCacheRepository() IRedisCacheRepository {
	return &RedisCacheRepository{}
}

func (repository *RedisCacheRepository) GetSecretKey() string {
	return redis.JWT_SECRET_KEY
}

func (repository *RedisCacheRepository) GetPreviousSecretKey() string {
	return redis.JWT_SECRET_KEY_PREVIOUS
}

func (repository *RedisCacheRepository) GetSecretKeyValue() (string, bool) {
	return redis.GetString(redis.JWT_SECRET_KEY)
}

func (repository *RedisCacheRepository) Put(key string, value interface{}, timeout time.Duration) error {
	return redis.Put(key, value, timeout)
}

func (repository *RedisCacheRepository) Get(key string) (string, bool) {
	return redis.GetString(key)
}

func (repository *RedisCacheRepository) Delete(keys []string) error {
	return redis.Delete(keys)
}

func (repository *RedisCacheRepository) IsExist(key string) bool {
	return redis.IsExist(key)
}

func (repository *RedisCacheRepository) GetStructData(key string, data interface{}) bool {
	return redis.GetStructData(key, &data)
}

func (repository *RedisCacheRepository) Lock(key string, options ...redsync.Option) bool {
	return redis.RedisLock(key, options...)
}

func (repository *RedisCacheRepository) Unlock(key string) bool {
	return redis.RedisUnlock(key)
}

func (repository *RedisCacheRepository) Scan(pattern string) ([]string, error) {
	return redis.Scan(pattern)
}

func (repository *RedisCacheRepository) HGet(key, field string) (string, error) {
	return redis.HGet(key, field)
}

func (repository *RedisCacheRepository) HSet(key, field, value string) error {
	return redis.HSet(key, field, value)
}

func (repository *RedisCacheRepository) HGetAll(key string) (map[string]string, error) {
	return redis.HGetAll(key)
}

func (repository *RedisCacheRepository) PutInHashMap(key string, values map[string]interface{}) error {
	return redis.PutInHashMap(key, values)
}

func (repository *RedisCacheRepository) GetHashMap(key string, field string) (result string, err error) {
	return redis.GetHashMap(key, field)
}

func (repository *RedisCacheRepository) DelHashMap(key string, field []string) (err error) {
	return redis.DelHashMap(key, field)
}
