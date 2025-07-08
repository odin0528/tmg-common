package redis

import (
	"fmt"
	"mgmt/common/utils"
	"strconv"
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

func GetScriptKeyWithModeAndFeatureType(gameName string, featureIndex, featureSecondIndex, mode, featureType int) string {
	return fmt.Sprintf("%s_%d_%d_%d_%d", gameName, featureIndex, featureSecondIndex, mode, featureType)
}

func GetGameCollectKey(gameName, playerAccount, currency string) string {
	return fmt.Sprintf("%s_%s_%s", gameName, playerAccount, currency)
}

func GetRoomRtpStatKey(roomId string, currency string) string {
	return GetCacheKey(KEY_ROOM_RTP_STATS, roomId, currency)
}

func GetScriptKey(gameName string, featureIndex, featureSecondIndex int) string {
	return fmt.Sprintf("%s_%d_%d", gameName, featureIndex, featureSecondIndex)
}

func GetLoginUsersCountByLoginTypeKey(date, loginType string) string {
	return GetCacheKey(LOGIN_TYPE_USER_HASH_KEY, date, loginType)
}

func GetAIAgentRoomInfoKey(gameName string) string {
	return GetCacheKey(AI_AGENT_ROOM_INFO_KEY, gameName)
}

func GetAILobbySettlementBetRecordKey() string {
	return GetCacheKey(AI_LOBBY_SETTLEMENT_BET_RECORD_KEY)
}

func GetAgentAccountByPlatformName(platformName string) string {
	return GetCacheKey(AGENT_ACCOUNT_BY_PLATFORM_NAME_KEY, platformName)
}

func GetGameListByAgent(agent string) string {
	return GetCacheKey(GAME_LIST_BY_AGENT_KEY, agent)
}

func GetDefaultGameOptionCacheKey(gameName, currency string, roomLevel int) string {
	roomLevelStr := strconv.Itoa(roomLevel)
	return GetCacheKey(GAME_OPTION_KEY, "default", gameName, currency, roomLevelStr)
}
