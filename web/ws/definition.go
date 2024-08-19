package ws

const (
	EVENT_WS_WEBSOCKET    = "websocket"
	EVENT_WS_SERVER_ERROR = "server_error"
	EVENT_WS_KEEP_ALIVE   = "keep_alive"
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
