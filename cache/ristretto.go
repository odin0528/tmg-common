package cache

import "github.com/dgraph-io/ristretto/v2"

func NewLocalCache[K ristretto.Key, V any](config *ristretto.Config[K, V]) *LocalCache[K, V] {
	cache, _ := ristretto.NewCache(config)
	return &LocalCache[K, V]{
		Cache: cache,
	}
}

func (lc *LocalCache[K, V]) Set(key K, value V) {
	lc.Cache.Set(key, value, 1)
	lc.Cache.Wait()
}

func (lc *LocalCache[K, V]) Get(key K) (V, bool) {
	return lc.Cache.Get(key)
}

func (lc *LocalCache[K, V]) Del(key K) {
	lc.Cache.Del(key)
}
