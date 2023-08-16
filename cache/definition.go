package cache

import "time"

type RedisCacheRepository struct {
}

type IRedisCacheRepository interface {
	GetSecretKeyValue() (string, bool)
	GetSecretKey() string
	GetPreviousSecretKey() string
	Put(key string, value interface{}, timeout time.Duration) error
	Get(key string) (string, bool)
	Delete(keys []string) error
	IsExist(key string) bool
	GetStructData(key string, data interface{}) bool
	Lock(key string) bool
	Unlock(key string) bool
	Scan(pattern string) ([]string, error)
	HGet(key, field string) (string, error)
	HSet(key, field, value string) error
	HGetAll(key string) (map[string]string, error)
}
