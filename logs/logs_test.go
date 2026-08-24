package logs

import (
	"testing"
	"mgmt/common/configs"
)

func TestWriteLog(t *testing.T) {
	configs.Init("../configs/env_example/")
	configs.LoadConfig([]string{"configs/env.ini"})

	InitLogs()

	Debug(LOG_TYPE_SYSTEM, "test", "system debug test", nil)
	Debug(LOG_TYPE_SYSTEM, "test", "system debug test 2", map[string]interface{}{
		"test_data1": "abc",
		"test_data2": 123,
	})

	Info(LOG_TYPE_SYSTEM, "test", "system info test", nil)
	Info(LOG_TYPE_SYSTEM, "test", "system info test 2", map[string]interface{}{
		"test_data3": "xyz",
		"test_data4": 987,
	})

	Error(LOG_TYPE_SYSTEM, "test", "system error test", nil)
	Error(LOG_TYPE_SYSTEM, "test", "system error test 2", map[string]interface{}{
		"test_data5": "qaz",
		"test_data6": 741,
	})

	Debug(LOG_TYPE_RECORD, "test", "record debug test", nil)
	Debug(LOG_TYPE_RECORD, "test", "record debug test 2", map[string]interface{}{
		"test_data1": "abc",
		"test_data2": 123,
	})

	Info(LOG_TYPE_RECORD, "test", "record info test", nil)
	Info(LOG_TYPE_RECORD, "test", "record info test 2", map[string]interface{}{
		"test_data3": "xyz",
		"test_data4": 987,
	})

	Error(LOG_TYPE_RECORD, "test", "record error test", nil)
	Error(LOG_TYPE_RECORD, "test", "record error test 2", map[string]interface{}{
		"test_data5": "qaz",
		"test_data6": 741,
	})

	type testBetRecord struct {
		RoundID  int    `json:"roundID" binding:"required"`
		RoomID   int    `json:"roomID" binding:"required"`
		GameName string `json:"gameName" binding:"required"`
	}

	betRecords := []testBetRecord{
		{
			RoundID:  123456789,
			RoomID:   123,
			GameName: "Test_Game",
		},
	}
	Record("test", betRecords)

	Debug(LOG_TYPE_CMS, "test", "cms debug test", nil)
	Debug(LOG_TYPE_CMS, "test", "cms debug test 2", map[string]interface{}{
		"test_data1": "abc",
		"test_data2": 123,
	})

	Info(LOG_TYPE_CMS, "test", "cms info test", nil)
	Info(LOG_TYPE_CMS, "test", "cms info test 2", map[string]interface{}{
		"test_data3": "xyz",
		"test_data4": 987,
	})

	Error(LOG_TYPE_CMS, "test", "cms error test", nil)
	Error(LOG_TYPE_CMS, "test", "cms error test 2", map[string]interface{}{
		"test_data5": "qaz",
		"test_data6": 741,
	})
}
