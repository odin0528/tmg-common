package response

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func GetCustomFormatError(str string, args ...interface{}) error {
	count := strings.Count(str, "%s")
	if count > len(args) {
		return errors.New("unexpected string:" + str)
	}
	return fmt.Errorf(str, args[:count]...)
}

func GetSuccessResponse() (int, Response) {
	return http.StatusOK, Response{
		Code: CODE_SUCCESS,
		Msg:  GetMsg(CODE_SUCCESS),
	}
}

func GetSuccessResponseWithData(data interface{}) (int, Response) {
	return http.StatusOK, Response{
		Code: CODE_SUCCESS,
		Msg:  GetMsg(CODE_SUCCESS),
		Data: data,
	}
}

func GetResponse(errCode int) (int, Response) {
	return http.StatusOK, Response{
		Code: errCode,
		Msg:  GetMsg(errCode),
	}
}

func GetMsg(code int) string {
	if msg, ok := API_CODE_MSG_MAP[code]; ok {
		return msg
	}

	return API_CODE_MSG_MAP[ERROR]
}
