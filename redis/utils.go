package redis

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"xxx/common/configs"
	"xxx/common/utils"
)

func GetPlayerBetInfoHashKey(account string) string {
	return fmt.Sprintf("%s_%s", BET_KEY, account)
}

func GetCacheKey(keys ...string) string {
	var builder strings.Builder
	for i, key := range keys {
		builder.WriteString(key)
		if i != len(keys)-1 {
			builder.WriteString(":")
		}
	}
	return builder.String()
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

func GetGameOptionCacheKey(platformName, gameName, currency string, roomLevel int) string {
	roomLevelStr := strconv.Itoa(roomLevel)
	return GetCacheKey(GAME_OPTION_KEY, platformName, gameName, currency, roomLevelStr)
}

func GetPlatformCurrencyGameOptionCacheKey(platformName, currency string) string {
	return GetCacheKey(GAME_OPTION_KEY, platformName, currency)
}

func IsNamespaceMaintain() bool {
	namespace := configs.Get(configs.SECTION_SYSTEM, configs.SYSTEM_NAMESPACE, "blue")
	return IsExistByAllType(GetCacheKey(NAMESAPCE_IS_MAINTAIN_KEY, namespace))
}

func GetAgentReverseByReverseCodeRedisKey(reverseCode string) string {
	return GetCacheKey(AGENT_REVERSE_PREFIX, REDIS_KEY_REVERSE_CODE, reverseCode)
}

func GetAgentReversePlayerTransactionKeyRedisKey(agentAccount, playerAccount string) string {
	return GetCacheKey(AGENT_REVERSE_PREFIX, agentAccount, REDIS_KEY_REVERSE_TRANSACTION_PLAYER, playerAccount)
}

func GetAgentGameRtpSettingHashKey(agentAccount, gameName string, controlType int) string {
	return GetCacheKey(agentAccount, gameName, strconv.Itoa(controlType))
}
