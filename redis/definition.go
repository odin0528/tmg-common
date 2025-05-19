package redis

import (
	"context"
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
	REGION_GAME_DAILY_CACHE   = "region_game_daily_cache"

	SINGLE_WALLET_SID_HASH_MAP_KEY                string = "single_wallet_sid_hash_key"
	SINGLE_WALLET_WITHDRAW_HASH_MAP_KEY           string = "single_wallet_withdraw_hash_key"
	SINGLE_WALLET_BET_RECORD_HASH_MAP_KEY         string = "sw_bet_record_hash_key"
	SINGLE_WALLET_BET_RECORD_RETRY_HASH_MAP_KEY   string = "sw_bet_record_retry_hash_key"
	SINGLE_WALLET_SYNC_BET_RECORD_HASH_MAP_KEY    string = "sw_sync_bet_record_hash_key"
	SINGLE_WALLET_AWARD_RECORD_HASH_MAP_KEY       string = "sw_award_record_hash_key"
	SINGLE_WALLET_AWARD_RECORD_RETRY_HASH_MAP_KEY string = "sw_award_record_retry_hash_key"

	SINGLE_WALLET_WITHDRAW_HASH_MAP_LOCK_KEY           string = "sw_witdraw_hash_lock"
	SINGLE_WALLET_BET_RECORD_HASH_MAP_LOCK_KEY         string = "sw_bet_record_hash_lock"
	SINGLE_WALLET_BET_RECORD_RETRY_HASH_MAP_LOCK_KEY   string = "sw_bet_record_retry_hash_lock"
	SINGLE_WALLET_SYNC_BET_RECORD_HASH_MAP_LOCK_KEY    string = "sw_sync_bet_record_hash_lock"
	SINGLE_WALLET_AWARD_RECORD_HASH_MAP_LOCK_KEY       string = "sw_award_record_hash_lock"
	SINGLE_WALLET_AWARD_RECORD_RETRY_HASH_MAP_LOCK_KEY string = "sw_award_record_retry_hash_lock"

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

	HOT_GAME_HASH_MAP_KEY string = "hot_game_hash_map_key"

	VERIFY_EXEC_SINGLE_WALLET_KEY = "verify_exec_single_wallet_key"

	INTERVAL_SECOND_CHECK_BET_RECORD_LASTET_UPDATED_AT string = "interval_second_check_bet_record_lastet_updated_at"

	REIDS_LOCK_DEFAULT_EXPIRE_SECOND                    = 30 * time.Second
	REDIS_LOCK_DEFAULT_RETRY_TIMES                  int = 30
	REDIS_LOCK_DEFAULT_RETRY_INTERVAL_TIME_DURATION     = 200 * time.Millisecond

	GAME_COLLECT_HASH_KEY               = "game_collect_hash_key"
	PLAYER_COLLECT_CHANGE_LIST_KEY      = "player_collect_change_list"
	PLAYER_COLLECT_CHANGE_LIST_LOCK_KEY = "player_collect_change_list_lock_key"

	PLAYER_UNBLOCK_HASH_KEY = "player_unblock_hash_key"

	DEFAULT_COLLECT_EXPIRED_TIME = time.Minute * 15

	LOGIN_TYPE_COUNT_HASH_KEY = "player_login_type_count_hash_key"
	LOGIN_TYPE_LOBBY          = "login_type_lobby"
	LOGIN_TYPE_GAME           = "login_type_game"

	AGENT_USE_API_HASH_KEY  = "agent_use_api_hash_key"
	SERVER_ADD_NEW_GAME_KEY = "server_add_new_game_key"

	RTP_NOTIFY_HASH_KEY = "rtp_notify_hash_key"

	WALLET_TYPE_AND_SYNC_MODE_KEY = "wallet_type_and_sync_mode"

	DO_DIRECT_BET_PREFIX_KEY = "do_direct_bet_prefix"

	GAME_OPTION_PREFIX_KEY = "game_option_prefix"
	COMMON_LIST_KEY        = "list"
)

const (
	ANNOUNCEMENT_KEY = "announcement"

	ANNOUNCEMENT_MODE_SYSTEM      int = 1
	ANNOUNCEMENT_MODE_NORMAL      int = 2
	ANNOUNCEMENT_MODE_REWARD      int = 3
	ANNOUNCEMENT_MODE_INTERACTION int = 4
	ANNOUNCEMENT_MODE_OFFICAL     int = 5

	ANNOUNCEMENT_FILTER_TYPE_ALL      int = 1
	ANNOUNCEMENT_FILTER_TYPE_AGENT    int = 2
	ANNOUNCEMENT_FILTER_TYPE_PLATFORM int = 3
	ANNOUNCEMENT_FILTER_TYPE_CURRENCY int = 4
)

const (
	MAILBOX_CLAIM_MAIL_LOCK_KEY = "mailbox_claim_mail_lock_key"
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
	Id               int         `json:"id"`
	Title            string      `json:"title"`
	Context          interface{} `json:"context"`
	Mode             int         `json:"mode"`
	TargetRule       string      `json:"target_rule"`
	AnnounceStartAt  *time.Time  `json:"announce_start_at"`
	AnnounceEndAt    *time.Time  `json:"announce_end_at"`
	RedirectGameName string      `json:"redirect_game_name"`
	Duration         int         `json:"duration"`
	NextAnnounceAt   time.Time   `json:"next_announce_at"`
	NotifyCount      int         `json:"notify_count"`
	IsDynamic        bool        `json:"is_dynamic"`

	ContextReplaceInfo interface{} `json:"context_replace_info"`
}

type TargetRule struct {
	Type int      `json:"type"`
	List []string `json:"list"`
}

type DeviceCountingFromLogin struct {
	TotalCount int
}

type UpdateHundredGameInfo struct {
	AgentAccountList []string `json:"agent_account_list"`
	CurrencyList     []string `json:"currency_list"`
}

type ContextReplaceInfo struct {
	TimeRangeText string  `json:"time_range_text"`
	PlayerAccount string  `json:"player_account"`
	GameName      string  `json:"game_name"`
	JackpotName   string  `json:"jackpot_name"`
	Award         float64 `json:"award"`
	Multiplier    float64 `json:"multiplier"`
}

type GameCollectInfo struct {
	Data        interface{} `json:"data"`
	ExpiredTime time.Time   `json:"expired_time"`
}

type Z struct {
	Score  float64
	Member string
}

// Pipeline 定義可用的 Redis pipeline 操作方法
type Pipeline interface {
	Set(key string, value interface{}, expiration time.Duration)
	Get(key string)
	Exec(ctx context.Context) ([]CmdResult, error)
	HSet(key string, field string, value interface{})
	HGet(key string, field string)

	// 可依需求增加更多常用方法
}

type TxPipeline interface {
	Pipeline
	Discard() error
}

// CmdResult 是每個 Redis 指令執行後的結果
type CmdResult struct {
	Result string
	Err    error
}

type PipelineWrapper struct {
	pipe redis.Pipeliner
}

type TxPipelineWrapper struct {
	PipelineWrapper
}
