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

// Package cache provide a Cache interface and some implement engine
// Usage:
//
// import(
//
//	"github.com/astaxie/beego/cache"
//
// )
//
// bm, err := cache.NewCache("memory", `{"interval":60}`)
//
// Use it like this:
//
//		bm.Put("astaxie", 1, 10 * time.Second)
//		bm.Get("astaxie")
//		bm.IsExist("astaxie")
//		bm.Delete("astaxie")
//
//	 more docs http://beego.me/docs/module/cache.md
package caches

import (
	"fmt"
	"time"

	"github.com/alicebob/miniredis"
)

// Cache interface contains all behaviors for cache adapter.
// usage:
//
//	cache.Register("file",cache.NewFileCache) // this operation is run in init method of file.go.
//	c,err := cache.NewCache("file","{....}")
//	c.Put("key",value, 3600 * time.Second)
//	v := c.Get("key")
//
//	c.Incr("counter")  // now is 1
//	c.Incr("counter")  // now is 2
//	count := c.Get("counter").(int)
type Cache interface {
	// Mock Redis
	Mock(s *miniredis.Miniredis)
	// get cached value by key.
	Get(key string) interface{}
	// GetMulti is a batch version of Get.
	GetMulti(keys []string) []interface{}
	// set cached value with key and expire time.
	Put(key string, val interface{}, timeout time.Duration) error
	// set cached value with key and no expire time.
	PutNoExpiry(key string, val interface{}) error
	// delete cached value by key.
	Delete(key string) error
	// increase cached int value by key, as a counter.
	Incr(key string) error
	// decrease cached int value by key, as a counter.
	Decr(key string) error
	// increase cached int value by key, as a counter.
	IncrFloat(key string, value float64) error
	// check if cached value exists or not.
	IsExist(key string) bool
	// clear all cache.
	ClearAll() error
	// start gc routine based on config string settings.
	StartAndGC(config string) error
	// get all keys
	Keys(filterText string, count int) []string
	// batch delete all keys
	DeleteMulti(keys []string) error
	// batch delete all keys without prefix
	DeleteMultiWithoutPrefix(keys []string) error
	// push data into tail of list
	LPush(key string, value interface{}) error
	// pop data from head of list
	RPop(key string) interface{}
	// return len of list
	LLen(key string) *int
	// it will keep only specified range of elements specified
	LTrim(key string, start, end int) error
	// it will return specified range of elements specified
	LRange(key string, start, end int) []interface{}
	// Set data into hash map for given values
	HMSet(key string, values map[string]interface{}) error
	// Get data from hash map for given field name
	HMGet(key string, fields []string) []string
	// Return all field and values from hashmap for given key pattern
	HGetAll(key string) map[string]string
	// DELETE fields in hashmap
	HDEL(key string, fields []string) bool
	// is exists fields in hashmap
	HEXISTS(key string, fields string) (bool, error)
	// HSET set fields in hashmap
	HSET(key string, fields string, value interface{}) error
	// HGET get fields in hashmap
	HGET(key string, fields string) (interface{}, error)
	// use Pipiline
	Pipeline([][]string) (interface{}, error)
}

// Instance is a function create a new Cache Instance
type Instance func() Cache

var adapters = make(map[string]Instance)

// Register makes a cache adapter available by the adapter name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, adapter Instance) {
	if adapter == nil {
		panic("cache: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("cache: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

// NewCache Create a new cache driver by adapter name and config string.
// config need to be correct JSON as string: {"interval":360}.
// it will start gc automatically.
func NewCache(adapterName, config string) (adapter Cache, err error) {
	instanceFunc, ok := adapters[adapterName]
	if !ok {
		err = fmt.Errorf("cache: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	adapter = instanceFunc()
	err = adapter.StartAndGC(config)
	if err != nil {
		adapter = nil
	}
	return
}

func NewMockCache(s *miniredis.Miniredis) (adapter Cache, err error) {
	instanceFunc, ok := adapters["redis"]
	if !ok {
		err = fmt.Errorf("cache: unknown adapter name redis (forgot to import?)")
		return
	}
	adapter = instanceFunc()
	adapter.Mock(s)
	return
}
