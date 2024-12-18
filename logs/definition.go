package logs

import (
	"go.uber.org/zap"
)

const (
	LOG_FILE_SYSTEM        = "system.log"
	LOG_FILE_RECORD        = "record.log"
	LOG_FILE_CMS           = "cms.log"
	LOG_FILE_PANIC_RECOVER = "panic_recover.log"
	LOG_FILE_BET           = "bet.log"
	LOG_FILE_HTTP          = "http.log"
	LOG_FILE_RISK_CONTROL  = "risk_control.log"
	LOG_FILE_GSI_API       = "gsi_api.log"
	LOG_FILE_RABBIT_MQ     = "rabbit_mq.log"

	TIME_KEY       = "created_at"
	LEVEL_KEY      = "level"
	NAME_KEY       = "logger"
	CALLER_KEY     = "from"
	MESSAGE_KEY    = "msg"
	STACKTRACE_KEY = "stacktrace"

	LOG_TYPE_ALL           = "all"
	LOG_TYPE_SYSTEM        = "system"
	LOG_TYPE_RECORD        = "record"
	LOG_TYPE_CMS           = "cms"
	LOG_TYPE_PANIC_RECOVER = "panic_recover"
	LOG_TYPE_BET           = "bet"
	LOG_TYPE_HTTP          = "http"
	LOG_TYPE_RISK_CONTROL  = "risk_control"
	LOG_TYPE_GSI_API       = "gsi_api"
	LOG_TYPE_RABBIT_MQ     = "rabbit_mq"

	SAVE_BET_RECORD = "Save bet record info"
	UNKNOWN_CALLER  = "Unknown caller"

	MAX_CALLER_COUNT  = 20
	BASE_SKIP_LAYER   = 1
	CALLER_SKIP_LAYER = 2
)

const (
	LOG_KEY_FUNC_CALLER = "func_caller"
	LOG_KEY_GAME_HUB    = "game_hub"
	LOG_KEY_MATCH_POOL  = "match_pool"
	LOG_KEY_BET_RECORD  = "bet_record"
	LOG_KEY_BET         = "bet"
	LOG_KEY_CAMP_WAR    = "camp_war_record"
	LOG_KEY_WATCHDOG    = "watchdog"

	LOG_KEY_CACHE               = "cache"
	LOG_KEY_BASE_ROOM           = "base_room"
	LOG_KEY_BASE_PLAYER         = "base_player"
	LOG_KEY_BATTLE_ROOM         = "battle_room"
	LOG_KEY_BATTLE_PLAYER       = "battle_player"
	LOG_KEY_BATTLE_ROOM_ADVANCE = "battle_room_advance"
	LOG_KEY_SLOT_BATTLE_ROOM    = "slot_battle_room"
	LOG_KEY_HUNDRED_ROOM        = "hundred_room"
	LOG_KEY_HUNDRED_PLAYER      = "hundred_player"
	LOG_KEY_SLOT_ROOM           = "slot_room"
	LOG_KEY_SLOT_PLAYER         = "slot_player"
	LOG_KEY_SINGLE_ROOM         = "single_room"
	LOG_KEY_SINGLE_PLAYER       = "single_player"
	LOG_KEY_MQ                  = "mq"
	LOG_KEY_AGENT               = "agent"
	LOG_KEY_CMS_USER            = "cms_user"
	LOG_KEY_API                 = "api"
	LOG_KEY_SINGLE_WALLET       = "single_wallet"
	LOG_KEY_ROLE                = "role"
	LOG_KEY_RISK_CONTROL        = "risk_control"
	LOG_KEY_MAILBOX             = "mailbox"
	LOG_KEY_COLLECT             = "collect"
	LOG_KEY_MQ_PRODUCER         = "producer"
	LOG_KEY_MQ_CONSUMER         = "consumer"

	LOG_KEY_CAMP_WAR_PLAYER_AWARD_RECORD = "camp_war_player_award_record"
)

const (
	FIELD_KEY_LOG_KEY           = "log_key"
	FIELD_KEY_FUNC_NAME         = "func_name"
	FIELD_KEY_FUNC_CALLER_STACK = "func_caller_stack"
	FIELD_KEY_PAYLOAD           = "data"

	FIELD_KEY_FSM_STATE            = "fsm_state"
	FIELD_KEY_SETTLEMENT           = "settlement"
	FIELD_KEY_BET_RECORD           = "bet_record"
	FIELD_KEY_CAMP_WAR_RECORD      = "camp_war_record"
	FIELD_KEY_PLATFORM_NAME        = "platform_name"
	FIELD_KEY_GAME_NAME            = "game_name"
	FIELD_KEY_PLAYER_NAME          = "player_name"
	FIELD_KEY_ROOM_ID              = "room_id"
	FIELD_KEY_ROUND_ID             = "round_id"
	FIELD_KEY_REDIST_KEY           = "redis_key"
	FIELD_KEY_START                = "start"
	FIELD_KEY_END                  = "end"
	FIELD_KEY_TIMEOUT              = "timeout"
	FIELD_KEY_CACHE_KEY            = "key"
	FIELD_KEY_CURRENCY             = "currency"
	FIELD_KEY_ROOM_LEVEL           = "room_level"
	FIELD_KEY_PLAYER               = "player"
	FIELD_IP                       = "ip"
	FIELD_KEY_BET_INFO             = "bet_info"
	FIELD_KEY_URL                  = "url"
	FIELD_KEY_HEADER               = "header"
	FIELD_KEY_BODY                 = "body"
	FIELD_KEY_SESSION_ID           = "session_id"
	FIELD_KEY_BALANCE              = "balance"
	FIELD_KEY_BET_ID               = "bet_id"
	FIELD_KEY_ERROR                = "error"
	FIELD_KEY_ERROR_CODE           = "error_code"
	FIELD_KEY_TIME_SECOND          = "time_second"
	FIELD_KEY_GAME_RESULT          = "game_result"
	FIELD_KEY_OLD_RESULT           = "old_result"
	FIELD_KEY_NEW_RESULT           = "new_result"
	FIELD_KEY_RISK_INFO            = "risk_info"
	FIELD_KEY_RETRY                = "retry"
	FIELD_KEY_CURRENT_NUMBER_INDEX = "current_number_index"
	FIELD_KEY_FEATURE_INDEX        = "feature_index"
	FIELD_KEY_FEATURE_SECOND_INDEX = "feature_second_index"
	FIELD_KEY_ODDS_TYPE            = "odds_type"
	FIELD_KEY_CONDITION            = "condition"
	FIELD_KEY_CAMP_POOL_MONEY      = "camp_pool_money"
	FIELD_KEY_UID                  = "uid"
)

var level zap.AtomicLevel
var systemLogger *zap.Logger
var recordLogger *zap.Logger
var cmsLogger *zap.Logger
var panicRecvoerLogger *zap.Logger
var betLogger *zap.Logger
var httpLogger *zap.Logger
var riskControlLogger *zap.Logger
var gsiApiLogger *zap.Logger
var rabbitMqLogger *zap.Logger

var systemLoggerCloseFunc func()
var recordLoggerCloseFunc func()
var cmsLoggerCloseFunc func()
var panicRecvoerLoggerCloseFunc func()
var betLoggerCloseFunc func()
var httpLoggerCloseFunc func()
var riskControlLoggerCloseFunc func()
var gsiApiLoggerCloseFunc func()
var rabbitMqLoggerCloseFunc func()

var currentDate string

type Field struct {
	Name string
	Data interface{}
}
