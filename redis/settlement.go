package redis

import (
	"encoding/json"
	"fmt"
	"mgmt/common/logs"

	"github.com/go-redis/redis/v8"
)

func GetUnsettleBetIds(account string) ([]string, error) {
	betIdsStr, err := GetHashMap(UNSETTLE_BET_ID_HASH_KEY, account)
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			return []string{}, nil
		}
		return []string{}, err
	}

	betIds := []string{}
	if err := json.Unmarshal([]byte(betIdsStr), &betIds); err != nil {
		logs.Error(logs.LOG_TYPE_SYSTEM, logs.LOG_KEY_BET_RECORD, fmt.Sprintf("Unmarshal err:%s", err.Error()), map[string]interface{}{
			logs.FIELD_KEY_PLAYER_NAME: account,
			logs.FIELD_KEY_PAYLOAD:     betIdsStr,
		})

		return []string{}, err
	}

	return betIds, nil
}

func GetAllUnsettleBetIds() (map[string][]string, error) {
	betIdsStringMap, err := GetAllHashMap(UNSETTLE_BET_ID_HASH_KEY)
	if err != nil && err.Error() != redis.Nil.Error() || len(betIdsStringMap) == 0 {
		return map[string][]string{}, err
	}

	accountBetIdsMap := map[string][]string{}
	for account, betIdsStr := range betIdsStringMap {
		betIds := []string{}

		err := json.Unmarshal([]byte(betIdsStr), &betIds)
		if err != nil {
			logs.Error(logs.LOG_TYPE_SYSTEM, logs.LOG_KEY_BET_RECORD, fmt.Sprintf("Unmarshal err:%s", err.Error()), map[string]interface{}{
				logs.FIELD_KEY_PLAYER_NAME: account,
				logs.FIELD_KEY_PAYLOAD:     betIdsStr,
			})

			continue
		}

		accountBetIdsMap[account] = betIds
	}

	return accountBetIdsMap, nil
}
