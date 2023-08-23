package ws

import (
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

	encode := base64Encode(byteArray)

	return encode
}

func GetServerErrorResponse(code int, msg string) []byte {
	return GetEventResponse(EVENT_WS_SERVER_ERROR, code, msg, nil)
}

func ParseEvent(message []byte) (Event, error) {
	var event Event

	decode := base64Decode(message)

	err := json.Unmarshal(decode, &event)

	return event, err
}

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
