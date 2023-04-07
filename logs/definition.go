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

	ALL    = "all"
	PANIC  = "panic"
	SYSTEM = "system"
	RECORD = "record"
	CMS    = "cms"

	SAVE_BET_RECORD = "Save bet record info"
	UNKNOWN_CALLER  = "Unknown caller"

	MAX_CALLER_COUNT  = 20
	BASE_SKIP_LAYER   = 1
	CALLER_SKIP_LAYER = 2
)

const (
	FUNC_CALLER = "func_caller"

	BASE_ROOM           = "base_room"
	BASE_PLAYER         = "base_player"
	BATTLE_ROOM         = "battle_room"
	BATTLE_PLAYER       = "battle_player"
	BATTLE_ROOM_ADVANCE = "battle_room_advance"
	HUNDRED_ROOM        = "hundred_room"
	HUNDRED_PLAYER      = "hundred_player"
	SLOT_ROOM           = "slot_room"
	SLOT_PLAYER         = "slot_player"
)

const (
	FIELD_KEY_LOG_KEY           = "log_key"
	FIELD_KEY_FUNC_NAME         = "func_name"
	FIELD_KEY_FUNC_CALLER_STACK = "func_caller_stack"
	FIELD_KEY_PAYLOAD           = "data"

	FIELD_KEY_FSM_STATE   = "fsm_state"
	FIELD_KEY_BET_RECORD  = "bet_record"
	FIELD_KEY_GAME_NAME   = "game_name"
	FIELD_KEY_PLAYER_NAME = "player_name"
	FIELD_KEY_ROOM_ID     = "room_id"
	FIELD_KEY_ROUND_ID    = "round_id"
)

var level zap.AtomicLevel
var systemLogger *zap.Logger
var recordLogger *zap.Logger
var cmsLogger *zap.Logger

type Field struct {
	Name string
	Data interface{}
}
