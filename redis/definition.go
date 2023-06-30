package redis

import "github.com/go-redis/redis/v8"

const (
	WOW_GAMING_MUTEX_PREFIX     string = "wow_gaming_mutex"
	AI_LIVE_CASINO_MUTEX_PREFIX string = "ai_live_casino_mutex"
	API_CENTER_MUTEX_PREFIX     string = "api_center_mutex"
	COMMON_MUTEX_PREFIX         string = "common_mutex"

	WOW_GAMING_PREFIX     string = "wow_gaming"
	AI_LIVE_CASINO_PREFIX string = "ai_live_casino"
	API_CENTER_PREFIX     string = "api_center"
	COMMON_PREFIX         string = "common"

	BET_RECORD_KEY           string = "bet_record"
	BET_KEY                  string = "bet"
	BET_PLAYER_LIST_HASH_KEY string = "bet_player_list_hash"
	UNSETTLE_BET_KEY         string = "unsettle_bet"
	UNSETTLE_BET_ID_HASH_KEY string = "unsettle_bet_id_hash"
	UNSTORE_SETTLEMENT_KEY   string = "unstore_settlement"

	MQ_BET_KEY string = "mq_bet"

	MUTEX_DURATION_SECOND = 20

	JWT_SECRET_KEY          string = "jwt_secret_key"
	JWT_SECRET_KEY_PREVIOUS string = "jwt_secret_key_previous"
	JWT_ACCOUNT_TOKEN       string = "jwt_account_token_"
	JWT_TOKEN_ACCOUNT       string = "jwt_token_account_"

	IS_PLAYING string = "is_playing"

	DEFAULT_SCAN_AMOUNT = 100
)

var REDIS_IS_NIL_ERR error = redis.Nil
