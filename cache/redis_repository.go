package cache

import (
	"time"
	"xxx/common/redis"
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

func (repository *RedisCacheRepository) Lock(key string) bool {
	return redis.RedisLock(key)
}

func (repository *RedisCacheRepository) Unlock(key string) bool {
	return redis.RedisUnlock(key)
}
