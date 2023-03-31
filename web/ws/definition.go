package ws

const (
	EVENT_WS_WEBSOCKET    = "websocket"
	EVENT_WS_SERVER_ERROR = "server_error"

	EVENT_WS_INIT_INFO    = "init_info"
	EVENT_WS_RECOVER_INFO = "recover_info"
	EVENT_WS_STATE_INFO   = "state_info"
	EVENT_WS_ROUND_INFO   = "round_info"
	EVENT_WS_RESULT_INFO  = "result_info"
	EVENT_WS_PAY_INFO     = "pay_info"
	EVENT_WS_TIME_INFO    = "time_info"
	EVENT_WS_GAME_OVER    = "game_over"
	EVENT_WS_READY        = "ready"

	EVENT_WS_BET   = "bet"
	EVENT_WS_REBET = "rebet"

	KEY_EVENT         = "event"
	KEY_DATA          = "data"
	KEY_CURRENT_STATE = "cur_state"
	KEY_CARD_SET_INFO = "card_set_info"
)

type Event struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data,omitempty"`
}

type EventResponse struct {
	Event string      `json:"event"`
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
}
