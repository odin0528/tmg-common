package logs

import (
	"go.uber.org/zap"
)

const (
	LOG_FILE_SYSTEM = "system.log"
	LOG_FILE_RECORD = "record.log"
	LOG_FILE_CMS    = "cms.log"

	TIME_KEY       = "created_at"
	LEVEL_KEY      = "level"
	NAME_KEY       = "logger"
	CALLER_KEY     = "from"
	MESSAGE_KEY    = "msg"
	STACKTRACE_KEY = "stacktrace"

	LOG_TYPE_ALL    = "all"
	LOG_TYPE_PANIC  = "panic"
	LOG_TYPE_SYSTEM = "system"
	LOG_TYPE_RECORD = "record"
	LOG_TYPE_CMS    = "cms"

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

	LOG_KEY_CACHE               = "cache"
	LOG_KEY_BASE_ROOM           = "base_room"
	LOG_KEY_BASE_PLAYER         = "base_player"
	LOG_KEY_BATTLE_ROOM         = "battle_room"
	LOG_KEY_BATTLE_PLAYER       = "battle_player"
	LOG_KEY_BATTLE_ROOM_ADVANCE = "battle_room_advance"
	LOG_KEY_HUNDRED_ROOM        = "hundred_room"
	LOG_KEY_HUNDRED_PLAYER      = "hundred_player"
	LOG_KEY_SLOT_ROOM           = "slot_room"
	LOG_KEY_SLOT_PLAYER         = "slot_player"
	LOG_KEY_SINGLE_ROOM         = "single_room"
	LOG_KEY_SINGLE_PLAYER       = "single_player"
	LOG_KEY_MQ                  = "mq"
)

const (
	FIELD_KEY_LOG_KEY           = "log_key"
	FIELD_KEY_FUNC_NAME         = "func_name"
	FIELD_KEY_FUNC_CALLER_STACK = "func_caller_stack"
	FIELD_KEY_PAYLOAD           = "data"

	FIELD_KEY_FSM_STATE     = "fsm_state"
	FIELD_KEY_SETTLEMENT    = "settlement"
	FIELD_KEY_BET_RECORD    = "bet_record"
	FIELD_KEY_PLATFORM_NAME = "platform_name"
	FIELD_KEY_GAME_NAME     = "game_name"
	FIELD_KEY_PLAYER_NAME   = "player_name"
	FIELD_KEY_ROOM_ID       = "room_id"
	FIELD_KEY_ROUND_ID      = "round_id"
	FIELD_KEY_REDIST_KEY    = "redis_key"
	FIELD_KEY_START         = "start"
	FIELD_KEY_END           = "end"
	FIELD_KEY_TIMEOUT       = "timeout"
	FIELD_KEY_CACHE_KEY     = "key"
	FIELD_KEY_CURRENCY      = "currency"
	FIELD_KEY_ROOM_LEVEL    = "room_level"
	FIELD_KEY_PLAYER        = "player"
)

var level zap.AtomicLevel
var systemLogger *zap.Logger
var recordLogger *zap.Logger
var cmsLogger *zap.Logger

type Field struct {
	Name string
	Data interface{}
}
