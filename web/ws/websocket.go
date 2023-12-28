package ws

import (
	"XXX/common/configs"
	"encoding/base64"
	"encoding/json"
)

func GetEventResponse(event string, code int, msg string, data interface{}) []byte {
	rsp := EventResponse{
		Event: event,
		Msg:   msg,
		Code:  code,
		Data:  data,
	}

	byteArray, _ := json.Marshal(rsp)

	if configs.Get(configs.SECTION_SYSTEM, configs.SYSTEM_ENABLE_WEBSOCKET_ENCODE, configs.YES) == configs.YES {
		encode := EncodeBybase64(byteArray)
		return encode
	}

	return byteArray
}

func GetServerErrorResponse(code int, msg string) []byte {
	return GetEventResponse(EVENT_WS_SERVER_ERROR, code, msg, nil)
}

func ParseEvent(message []byte) (Event, error) {
	var event Event

	if configs.Get(configs.SECTION_SYSTEM, configs.SYSTEM_ENABLE_WEBSOCKET_DECODE, configs.YES) == configs.YES {
		decode, _ := DecodeByBase64(message)
		err := json.Unmarshal(decode, &event)
		return event, err
	}

	err := json.Unmarshal(message, &event)

	return event, err
}

func EncodeBybase64(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func DecodeByBase64(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
