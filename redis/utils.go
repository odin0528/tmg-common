package redis

import "fmt"

func GetPlayerBetHashKey(account string) string {
	return fmt.Sprintf("%s_%s", BET_KEY, account)
}

func GetUnsettleBetHashKey(account string) string {
	return fmt.Sprintf("%s_%s", BET_KEY, account)
}
