package cache

import (
	"xxx/common/redis"
	"time"
)

func NewCacheRepository() ICacheRepository {
	return &CacheRepository{}
}

func (repository *CacheRepository) GetSecretKey() string {
	return redis.JWT_SECRET_KEY
}

func (repository *CacheRepository) GetPreviousSecretKey() string {
	return redis.JWT_SECRET_KEY_PREVIOUS
}

func (repository *CacheRepository) GetSecretKeyValue() (string, bool) {
	return redis.GetString(redis.JWT_SECRET_KEY)
}

func (repository *CacheRepository) Put(key string, value interface{}, timeout time.Duration) error {
	return redis.Put(key, value, timeout)
}

func (repository *CacheRepository) Get(key string) (string, bool) {
	return redis.GetString(key)
}

func (repository *CacheRepository) Delete(keys []string) error {
	return redis.Delete(keys)
}

func (repository *CacheRepository) IsExist(key string) bool {
	return redis.IsExist(key)
}

func (repository *CacheRepository) GetStructData(key string, data interface{}) bool {
	return redis.GetStructData(key, &data)
}

func (repository *CacheRepository) Lock(key string) bool {
	return redis.RedisLock(key)
}

func (repository *CacheRepository) Unlock(key string) bool {
	return redis.RedisUnlock(key)
}