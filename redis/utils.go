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
