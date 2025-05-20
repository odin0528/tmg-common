package redis

import (
	"fmt"
	"mgmt/common/utils"
	"mgmt/pkg/model"
	"strings"
	"time"
)

func GetPlayerBetInfoHashKey(account string) string {
	return fmt.Sprintf("%s_%s", BET_KEY, account)
}

func GetCacheKey(keys ...string) string {
	result := ""
	for i, key := range keys {
		result += key

		if i != (len(keys) - 1) {
			result += ":"
		}
	}

	return result
}

func GetAllCacheKey(keys ...string) string {
	result := ""
	for i, key := range keys {
		result += key

		if i != (len(keys) - 1) {
			result += ":"
		}
	}

	return result + ":*"
}

func GetSingleWalletFieldKey(roundId, playerAccount string) string {
	return fmt.Sprintf("%s_%s", roundId, playerAccount)
}

func GetSingleWalletLockKey(hashMapKey, filedKey string) string {
	return fmt.Sprintf("%s_%s", hashMapKey, filedKey)
}

func GetDailyBetRankKey(gameName, currency string) string {
	now := time.Now()
	date := utils.ConvertToDateString(now)

	cacheKey := GetCacheKey(DAILY_BET_RANK_KEY, date, gameName, currency)
	return cacheKey
}

func GetScriptKey(gameName, tableName string, oddsType int) string {
	return fmt.Sprintf("%s:%s:%s:%d",
		RISK_CONTROL_SCRIPT_KEY,
		strings.ToLower(gameName),
		tableName,
		oddsType,
	)
}

func GetScriptKeyWithCondition(gameName, tableName, condition string, oddsType int) string {
	return fmt.Sprintf("%s:%s:%s:%s:%d",
		RISK_CONTROL_SCRIPT_KEY,
		strings.ToLower(gameName),
		tableName,
		condition,
		oddsType,
	)
}

func GetScriptKeyByGame(gameName, tableName, condition string, featureIndex, oddsType int) string {
	switch strings.ToUpper(gameName) {
	//bonus也有condition的話跑第一個
	case model.GAME_NAME_BIG_FIVE_GAME:
		return GetScriptKeyWithCondition(gameName, tableName, condition, oddsType)
	default:
		if featureIndex == 0 {
			return GetScriptKeyWithCondition(gameName, tableName, condition, oddsType)
		} else {
			return GetScriptKey(gameName, tableName, oddsType)
		}
	}
}

func GetScriptKeyWithModeAndFeatureType(gameName string, featureIndex, featureSecondIndex, mode, featureType int) string {
	return fmt.Sprintf("%s_%d_%d_%d_%d", gameName, featureIndex, featureSecondIndex, mode, featureType)
}

func GetGameCollectKey(gameName, playerAccount, currency string) string {
	return fmt.Sprintf("%s_%s_%s", gameName, playerAccount, currency)
}

func GetRoomRtpStatKey(roomId string, currency string) string {
	return GetCacheKey(KEY_ROOM_RTP_STATS, roomId, currency)
}
