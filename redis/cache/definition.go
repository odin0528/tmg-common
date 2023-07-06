package cache

import "time"

type CacheRepository struct {
}

type ICacheRepository interface {
	GetSecretKeyValue() (string, bool)
	GetSecretKey() string
	GetPreviousSecretKey() string
	Put(key string, value interface{}, timeout time.Duration) error
	Get(key string) (string, bool)
	Delete(keys []string) error
	IsExist(key string) bool
}
