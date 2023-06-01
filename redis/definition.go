package redis

const (
	WOW_GAMING_MUTEX_PREFIX     string = "wow_gaming_mutex:"
	AI_LIVE_CASINO_MUTEX_PREFIX string = "ai_live_casino_mutex:"
	API_CENTER_MUTEX_PREFIX     string = "wow_gaming_mutex:"

	BET_RECORD_KEY string = "bet_record"

	MUTEX_DURATION_SECOND = 20

	JWT_SECRET_KEY          string = "jwt_secret_key"
	JWT_SECRET_KEY_PREVIOUS string = "jwt_secret_key_previous"
	JWT_ACCOUNT_TOKEN       string = "jwt_account_token_"
	JWT_TOKEN_ACCOUNT       string = "jwt_token_account_"
)
