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

package redis

import (
	"fmt"
	"mgmt/common/caches"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
)

func setup() (caches.Cache, error) {
	return caches.NewCache("redis", `{"conn": "127.0.0.1:6379", "password": "123456"}`)
}
func TestRedisCache(t *testing.T) {
	bm, err := setup()

	if err != nil {
		t.Error("init err")
	}
	timeoutDuration := 10 * time.Second
	if err = bm.Put("astaxie", 1, timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("astaxie") {
		t.Error("check err")
	}

	time.Sleep(11 * time.Second)

	if bm.IsExist("astaxie") {
		t.Error("check err")
	}
	if err = bm.Put("astaxie", 1, timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if v, _ := redis.Int(bm.Get("astaxie"), err); v != 1 {
		t.Error("get err")
	}

	if err = bm.Incr("astaxie"); err != nil {
		t.Error("Incr Error", err)
	}

	if v, _ := redis.Int(bm.Get("astaxie"), err); v != 2 {
		t.Error("get err")
	}

	if err = bm.Decr("astaxie"); err != nil {
		t.Error("Decr Error", err)
	}

	if v, _ := redis.Int(bm.Get("astaxie"), err); v != 1 {
		t.Error("get err")
	}
	bm.Delete("astaxie")
	if bm.IsExist("astaxie") {
		t.Error("delete err")
	}

	//test string
	if err = bm.Put("astaxie", "author", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("astaxie") {
		t.Error("check err")
	}

	if v, _ := redis.String(bm.Get("astaxie"), err); v != "author" {
		t.Error("get err")
	}

	//test GetMulti
	if err = bm.Put("astaxie1", "author1", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("astaxie1") {
		t.Error("check err")
	}

	vv := bm.GetMulti([]string{"astaxie", "astaxie1"})
	if len(vv) != 2 {
		t.Error("GetMulti ERROR")
	}
	if v, _ := redis.String(vv[0], nil); v != "author" {
		t.Error("GetMulti ERROR")
	}
	if v, _ := redis.String(vv[1], nil); v != "author1" {
		t.Error("GetMulti ERROR")
	}

	// test clear all
	if err = bm.ClearAll(); err != nil {
		t.Error("clear all err")
	}
}

func TestListCached(t *testing.T) {
	cache, err := setup()
	if err != nil {
		t.Error("init err")
	}

	if err := cache.LPush("mylist", 10); nil != err {
		t.Error("LPush ERROR")
	}

	// Keep only one element
	if err := cache.LTrim("mylist", 0, 0); nil != err {
		t.Error("LTrim ERROR")
	}

	if val := cache.LLen("mylist"); nil != val && 1 != *val {
		t.Error("LLen ERROR")
	}

	if val := cache.LRange("mylist", 0, -1); nil == val {
		t.Error("LRange ERROR")
	}

	if val := cache.RPop("mylist"); nil == val {
		t.Error("RPop ERROR")
	}
}

func TestHashmapCached(t *testing.T) {
	key := "my_hashmap"
	cache, err := setup()
	if err != nil {
		t.Error("init err")
	}

	type metadata struct {
		Bet int `json:"bet"`
	}

	data := map[string]interface{}{}
	baseNum := 100000
	for i := 0; i < baseNum; i++ {
		k := fmt.Sprintf("ze_a%d", i)
		data[k] = metadata{Bet: i}
	}
	// log.Println(data)

	if err := cache.HMSet(key, data); nil != err {
		t.Error("HMSet ERROR")
	}

	data2 := map[string]interface{}{}
	extraNum := 100
	for i := 0; i < extraNum; i++ {
		k := fmt.Sprintf("guest_a%d", i)
		data2[k] = metadata{Bet: i}
	}
	if err := cache.HMSet(key, data2); nil != err {
		t.Error("HMSet ERROR")
	}

	players := cache.HMGet(key, []string{"ze_a1", "ze_a4"})
	if 2 != len(players) {
		t.Fatal(fmt.Sprintf("HMGet key number(%d) is not 2.", len(players)))
	}

	results := cache.HGetAll(key)
	if extraNum+baseNum != len(results) {
		t.Fatal(fmt.Sprintf("HGetAll number(%d) is not %d.", len(results), (extraNum + baseNum)))
	}

	success := cache.HDEL(key, []string{"ze_a3"})
	if true != success {
		t.Fatal("HDEL key failed.")
	}

	results = cache.HGetAll(key)
	if extraNum+baseNum-1 != len(results) {
		t.Fatal(fmt.Sprintf("HGetAll number(%d) is not %d.", len(results), (extraNum + baseNum - 1)))
	}

}
