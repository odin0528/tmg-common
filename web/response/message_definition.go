package response

const (
	MSG_SUCCESS = "success"

	MSG_NOT_SUPPORT_FUNCTION = "not support function"
)

// api
const (
	MSG_PARAMETER_ERROR               = "invalid parameters"
	MSG_API_CENTER_ERROR              = "api center response error"
	MSG_PARSE_API_RESPONSE_ERROR      = "parse api response error"
	MSG_PARSE_API_RESPONSE_DATA_ERROR = "parse api response data error"
)

// ws
const (
	MSG_WS_QUERY_PARAMETER_ERROR = "invalid ws query parameters"
	MSG_WS_IS_CLOSED             = "ws is closed"
)

// player
const (
	MSG_PLAYER_RELOGIN                 = "player relogin"
	MSG_PLAYER_RECOVER                 = "player is recover to the game"
	MSG_PLAYER_MONEY_NOT_ENOUGH        = "player's money is not enough"
	MSG_PLAYER_MOENY_IS_OUT_OF_MAX_BET = "player's bet money is out of max bet"
	MSG_PLAYER_CURRENT_ROUND_HAS_BET   = "player current round has bet"
	MSG_PLAYER_LAST_ROUND_NOT_BET      = "player last round not bet"
	MSG_PLAYER_IS_IN_ANOTHER_GAME      = "player is in another game: %s"
	MSG_PLAYER_ACTION_IN_WRONG_STATE   = "player can't use this action in the state"
	MSG_PLAYER_REQUEST_TOO_FREQUENTLY  = "player request too frequently"
	MSG_PLAYER_ACTOION_ALREADY_DONE    = "player's action already done"
	MSG_PLAYER_ACTOION_IS_NOT_ALLOW    = "player's action is NOT allow"
	MSG_PLAYER_EVENT_IS_NOT_SUPPOERED  = "event is not supported"
)

// game
const (
	MSG_ROOM_IS_FINISH   = "room is finish"
	MSG_GAME_IS_NOT_INIT = "game is not init"
	MSG_NOT_IN_STATE     = "not in %s state, current state: %s"
)

// internal
const (
	MSG_MARSHAL_ERROR      = "marshal error"
	MSG_UNMARSHAL_ERROR    = "unmarshal error"
	MSG_CONVERT_TYPE_ERROR = "convert type error"
	MSG_DATA_NOT_FOUND     = "%s doesn't exist %s"
)
