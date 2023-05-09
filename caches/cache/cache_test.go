package cache

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
	"game_server/common/configs"
)

func TestBatchRedisLock(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	keys := []string{}
	m := map[int]bool{}
	for i := 0; i < 1000; i++ {
		v := rand.Intn(100)
		if exist, _ := m[v]; true == exist {
			continue
		}

		keys = append(keys, strconv.Itoa(v))
		m[v] = true
	}

	success := BatchRedisLock(keys)
	log.Println("result: ", success)
}

func TestPipeline(t *testing.T) {
	configs.Init("../../configs/example/")
	ConfigNewCache()

	testSlice := [][]string{
		{"SET", "testKey", "testValue"},
		{"GET", "testKey"},
		{"SET", "testKey2", "testValue2"},
		{"GET", "testKey2"},
		{"LPUSH", "testList", "testListValue"},
		{"LPUSH", "testList", "testListValue"},
		{"LPUSH", "testList", "testListValue"},
		{"RPOP", "testList"},
		{"RPOP", "testList"},
	}

	replys, err := Pipeline(testSlice)
	if err != nil {
		fmt.Println(err)
	}

	for _, r := range replys {
		fmt.Printf("%s\n", r)
	}
}
