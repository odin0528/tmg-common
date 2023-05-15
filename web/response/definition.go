package response

var API_CODE_MSG_MAP = map[int]string{
	ERROR: "Internal server error.",

	CODE_SUCCESS:               "Success",
	CODE_API_INVALID_PARAMETER: "invalid parameter",
	CODE_API_ILLEGAL_PARAMETER: "illegal parameter",

	CODE_DATABASE_ABNORMAL: "database abnormal",

	CODE_AUTH_API_KEY_FAILED: "auth api key failed",

	CODE_CACHE_KEY_COLLISION: "cache key collision",
	CODE_CACHE_PUT_ERROR:     "fail to put cache",
	CODE_CACHE_GET_ERROR:     "fail to get cache",
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
