package cache

import (
	"time"

	"github.com/dgraph-io/ristretto/v2"
	"github.com/go-redsync/redsync/v4"
)

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
	Lock(key string, options ...redsync.Option) bool
	Unlock(key string) bool
	Scan(pattern string) ([]string, error)
	HGet(key, field string) (string, error)
	HSet(key, field, value string) error
	HGetAll(key string) (map[string]string, error)

	PutInHashMap(key string, values map[string]interface{}) error
	GetHashMap(key string, field string) (result string, err error)
	DelHashMap(key string, field []string) (err error)
}

type LocalCache[K ristretto.Key, V any] struct {
	Cache *ristretto.Cache[K, V]
}
