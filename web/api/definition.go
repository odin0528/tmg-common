package api

type HTTP_METHOD string
type WebAPICallback func([]byte, interface{}, error)

const (
	DEFAULT_API_TIME_OUT = 30

	METHOD_GET    HTTP_METHOD = "GET"
	METHOD_POST   HTTP_METHOD = "POST"
	METHOD_DELETE HTTP_METHOD = "DELETE"
	METHOD_PUT    HTTP_METHOD = "PUT"

	HEADER_AUTHORIZATION = "authorization"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
