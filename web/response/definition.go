package response

var API_CODE_MSG_MAP = map[int]string{
	ERROR: "Internal server error.",

	CODE_SUCCESS:               "Success",
	CODE_API_INVALID_PARAMETER: "invalid parameter",
	CODE_API_ILLEGAL_PARAMETER: "illegal parameter",

	CODE_DATABASE_ABNORMAL: "database abnormal",
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
