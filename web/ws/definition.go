package ws

const (
	EVENT_WS_WEBSOCKET      = "websocket"
	EVENT_WS_SERVER_ERROR   = "server_error"
	EVENT_WS_KEEP_ALIVE     = "keep_alive"
	EVENT_WS_ANNOUNCEMENT   = "announcement"
	EVENT_WS_UPDATE_BALANCE = "update_balance"

	HEX_MAX_BIT = 255
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
