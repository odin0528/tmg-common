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

func GetErrorCodeResponse(errCode int) Response {
	return Response{
		Code: errCode,
		Msg:  GetMsg(errCode),
	}
}

func GetResponse(errCode int, msg string) Response {
	return Response{
		Code: errCode,
		Msg:  msg,
	}
}

func GetMsg(code int) string {
	if msg, ok := API_CODE_MSG_MAP[code]; ok {
		return msg
	}

	return API_CODE_MSG_MAP[CODE_INTERNAL_ERROR]
}
