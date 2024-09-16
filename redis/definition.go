package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
)

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
	BET_COUNT_HASH_MAP_KEY   string = "bet_count_hash_map_key"
	BET_PLAYER_LIST_HASH_KEY string = "bet_player_list_hash"
	UNSETTLE_BET_KEY         string = "unsettle_bet"
	UNSETTLE_BET_ID_HASH_KEY string = "unsettle_bet_id_hash"
	UNSTORE_SETTLEMENT_KEY   string = "unstore_settlement"

	MQ_BET_KEY string = "mq_bet"

	MUTEX_DURATION_SECOND = 20

	JWT_SECRET_KEY          string = "jwt_secret_key"
	JWT_SECRET_KEY_PREVIOUS string = "jwt_secret_key_previous"
	JWT_ACCOUNT_TOKEN       string = "jwt_account_token_"
	JWT_CMS_ACCOUNT_TOKEN   string = "jwt_cms_account_token_"
	JWT_TOKEN_ACCOUNT       string = "jwt_token_account_"

	IS_PLAYING       string = "is_playing"
	IS_PLAYING_COUNT string = "is_playing_count"
	FOCUS_GAME       string = "focus_game"

	AGENT_ID_SERIAL_NUMBER_PREFIX string = "agent_id_serial_num_"

	LOGIN_DETAIL_HASH_KEY string = "cms_login_hash_key"

	LOGIN_ACCOUNT_HASH_KEY string = "cms_login_account_hash_key"

	PLAYER_TOTAL_PROFIT_CACHE = "player_total_profit_cache"
	PLAYER_GAME_DAILY_CACHE   = "player_game_daily_cache"
	AGENT_GAME_DAILY_CACHE    = "agent_game_daily_cache"

	SINGLE_WALLET_SID_HASH_MAP_KEY              string = "single_wallet_sid_hash_key"
	SINGLE_WALLET_WITHDRAW_HASH_MAP_KEY         string = "single_wallet_withdraw_hash_key"
	SINGLE_WALLET_BET_RECORD_HASH_MAP_KEY       string = "sw_bet_record_hash_key"
	SINGLE_WALLET_BET_RECORD_RETRY_HASH_MAP_KEY string = "sw_bet_record_retry_hash_key"
	SINGLE_WALLET_SYNC_BET_RECORD_HASH_MAP_KEY  string = "sw_sync_bet_record_hash_key"

	SINGLE_WALLET_WITHDRAW_HASH_MAP_LOCK_KEY         string = "sw_witdraw_hash_lock"
	SINGLE_WALLET_BET_RECORD_HASH_MAP_LOCK_KEY       string = "sw_bet_record_hash_lock"
	SINGLE_WALLET_BET_RECORD_RETRY_HASH_MAP_LOCK_KEY string = "sw_bet_record_retry_hash_lock"
	SINGLE_WALLET_SYNC_BET_RECORD_HASH_MAP_LOCK_KEY  string = "sw_sync_bet_record_hash_lock"

	SINGLE_WALLET_TRANSFER_RECORD_QUERY string = "sw_record_query"

	SINGLE_WALLET_TRANSFER_RECORD_QUERY_CACHE_DURATION = time.Second * 30

	IS_UPDATE_WHITE_LIST = "is_update_white_list"

	DAILY_BET_RANK_KEY = "daily_bet_rank_key"

	DAILY_BET_RANK_KEY_DURATION = time.Minute * 1

	DEFAULT_SCAN_AMOUNT = 100

	DEBUG_MAX_SINGLE_WALLET_WITHDRAW_RETRY_CACHE_KEY = "max_single_wallet_withdraw_retry_limit"

	RISK_CONTROL_SCRIPT_WEIGHT_HASH_KEY      = "risk_control_script_weight_hash_key"
	RISK_CONTROL_SCRIPT_WEIGHT_HASH_LOCK_KEY = "risk_control_script_weight_hash_lock"
	RISK_CONTROL_ODDS_TYPE_BACKUP_HASH_KEY   = "risk_control_odds_type_backup_hash_key"
	RISK_CONTROL_HASH_KEY                    = "risk_control_hash_key"

	GAME_CURRENT_RTP_KEY           = "game_current_rtp"
	GAME_CURRENT_RTP_DATE_LIST_KEY = "game_current_rtp_date_list"
	GAME_COUNTING_BUY_FEATURE_KEY  = "game_counting_buy_feature"

	UPDATE_HUNDRED_GAME_KEY = "is_update_hundred_game"

	DEVICE_COUNTING_FROM_LOGIN = "device_counting_from_login"
)

const (
	ANNOUNCEMENT_KEY = "announcement"

	ANNOUNCEMENT_MODE_SYSTEM      int = 1
	ANNOUNCEMENT_MODE_NORMAL      int = 2
	ANNOUNCEMENT_MODE_REWARD      int = 3
	ANNOUNCEMENT_MODE_INTERACTION int = 4
	ANNOUNCEMENT_MODE_OFFICAL     int = 5
)

var ANNOUNCEMENT_MODE_LIST = []int{
	ANNOUNCEMENT_MODE_SYSTEM,
	ANNOUNCEMENT_MODE_NORMAL,
	ANNOUNCEMENT_MODE_REWARD,
	ANNOUNCEMENT_MODE_INTERACTION,
	ANNOUNCEMENT_MODE_OFFICAL,
}

var REDIS_IS_NIL_ERR error = redis.Nil

type CmsLoginTimeoutCacheInfo struct {
	ExpiredTime time.Time `json:"expired_time"`
}

type GameRtpInfo struct {
	CurrentRtp   decimal.Decimal
	LastAlertRtp decimal.Decimal
	TotalIncome  decimal.Decimal
	TotalBet     decimal.Decimal
	TotalCount   int
}

type GameCountingBuyFeatureInfo struct {
	FGCount int
	BGCount int
}

type Announcement struct {
	Id               int        `json:"id"`
	Title            string     `json:"title"`
	Context          string     `json:"context"`
	Mode             int        `json:"mode"`
	TargetRule       string     `json:"target_rule"`
	AnnounceStartAt  *time.Time `json:"announce_start_at"`
	AnnounceEndAt    *time.Time `json:"announce_end_at"`
	RedirectGameName string     `json:"redirect_game_name"`
	Duration         int        `json:"duration"`
	NextAnnounceAt   time.Time  `json:"next_announce_at"`
	NotifyCount      int        `json:"notify_count"`
}

type DeviceCountingFromLogin struct {
	TotalCount int
}

type UpdateHundredGameInfo struct {
	AgentAccountList []string `json:"agent_account_list"`
	CurrencyList     []string `json:"currency_list"`
}
