package redis

import (
	"fmt"
	"mgmt/common/utils"
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

func GetScriptKey(gameName string, featureIndex, featureSecondIndex int) string {
	return fmt.Sprintf("%s_%d_%d", gameName, featureIndex, featureSecondIndex)
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
