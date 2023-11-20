package response

import (
	"errors"
	"fmt"
	"strings"
)

func GetCustomFormatError(str string, args ...interface{}) error {
	count := strings.Count(str, "%s")
	if count > len(args) {
		return errors.New("unexpected string:" + str)
	}
	return fmt.Errorf(str, args[:count]...)
}

func GetSuccessResponse() Response {
	return Response{
		Code: CODE_SUCCESS,
		Msg:  GetMsg(CODE_SUCCESS),
	}
}

func GetSuccessResponseWithData(data interface{}) Response {
	return Response{
		Code: CODE_SUCCESS,
		Msg:  GetMsg(CODE_SUCCESS),
		Data: data,
	}
}

func GetErrorCodeResponse(errCode int, args ...interface{}) Response {
	return Response{
		Code: errCode,
		Msg:  fmt.Sprintf(GetMsg(errCode), args...),
	}
}

func GetResponse(errCode int, msg string) Response {
	return Response{
		Code: errCode,
		Msg:  msg,
	}
}

func GetResponseWithData(errCode int, msg string, data interface{}) Response {
	return Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	}
}

func GetMsg(code int) string {
	if msg, ok := API_CODE_MSG_MAP[code]; ok {
		return msg
	}

	return API_CODE_MSG_MAP[CODE_INTERNAL_ERROR]
}
