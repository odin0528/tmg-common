package redis

import (
	"fmt"
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
