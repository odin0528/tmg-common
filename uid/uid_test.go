package uid

import (
	"strconv"
	"sync"
	"testing"
	"time"

	"mgmt/common/configs"

	"github.com/stretchr/testify/assert"
)

func TestGetCustomizeUniqueIDFormat(t *testing.T) {
	configs.Init("../../common/configs/example/")

	Init()

	result, err := getCustomizeUniqueIDByID(int64(1658670915090649090))

	assert.Nil(t, err)

	assert.Equal(t, "1011100000100110010000110100011000101010000000000100000000001", strconv.FormatInt(result, BINARY))
}

func TestConcurrencyGenerateUniqueID(t *testing.T) {
	configs.Init("../../common/configs/example/")

	Init()

	tracedIDMap := map[int]bool{}
	var mutex sync.RWMutex
	testTimes := 1000000

	for i := 0; i < testTimes; i++ {
		go func(v int) {
			uniqueID := GenerateUniqueID()

			mutex.Lock()
			defer mutex.Unlock()

			_, ok := tracedIDMap[uniqueID]
			assert.False(t, ok)
			tracedIDMap[uniqueID] = true
		}(i)
	}

	time.Sleep(time.Second * 3)
}

func TestGenerateUniqueIDAlwaysLessOrEqualThanSnowFlakeID(t *testing.T) {
	configs.Init("../../common/configs/example/")

	Init()

	testTimes := 10000

	for i := 0; i < testTimes; i++ {
		uniqueID := GenerateUniqueID()

		compareID, _ := getCustomizeUniqueIDByID(uniqueIDNode.Generate().Int64())

		assert.True(t, (compareID >= int64(uniqueID)))

		bu := get64BitBinaryUniqueID(int64(uniqueID))

		if (i+1)%SEQUENCE_ID_THRESHOLD == 0 {
			v, _ := strconv.ParseInt(getCustomizeBinarySequenceID(bu), BINARY, UNIQUE_ID_LENGTH)
			assert.Equal(t, int64(0), v)
		}
	}
}
