package api

import "sync/atomic"

type HTTP_METHOD string
type WebAPICallback func([]byte, interface{}, error)

const (
	DEFAULT_API_TIME_OUT = 30
	BET_API_TIME_OUT     = 3

	GAME_SERVER_BET_TIME_OUT = BET_API_TIME_OUT + 1

	METHOD_GET    HTTP_METHOD = "GET"
	METHOD_POST   HTTP_METHOD = "POST"
	METHOD_DELETE HTTP_METHOD = "DELETE"
	METHOD_PUT    HTTP_METHOD = "PUT"

	HEADER_KEY_CONTENT_TYPE = "Content-Type"
	HEADER_KEY_AUTH         = "Authorization"
	HEADER_KEY_USER_AGENT   = "User-Agent"

	CONTENT_TYPE_JSON = "application/json;charset=UTF-8"

	DATA_MONEY_KEY = "money"

	REDIS_KEY_ENABLE_CUSTOM_USER_AGENT = "enable_custom_user_agent"
)

var enableCustomUserAgent atomic.Bool

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
