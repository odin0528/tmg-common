package response

var API_CODE_MSG_MAP = map[int]string{
	ERROR: "Internal server error.",

	CODE_SUCCESS:             "Success",
	CODE_API_PARAMETER_ERROR: "API parameter error.",
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
