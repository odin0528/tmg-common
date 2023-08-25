package ws

import (
	"encoding/base64"
	"encoding/json"
	"game_server/common/configs"
)

func GetEventResponse(event string, code int, msg string, data interface{}) []byte {
	rsp := EventResponse{
		Event: event,
		Msg:   msg,
		Code:  code,
		Data:  data,
	}

	byteArray, _ := json.Marshal(rsp)

	if configs.Get(configs.SECTION_SYSTEM, configs.SYSTEM_ENABLE_WEBSOCKET_ASE, configs.YES) == configs.YES {
		encode := base64Encode(byteArray)
		return encode
	}

	return byteArray
}

func GetServerErrorResponse(code int, msg string) []byte {
	return GetEventResponse(EVENT_WS_SERVER_ERROR, code, msg, nil)
}

func ParseEvent(message []byte) (Event, error) {
	var event Event

	err := json.Unmarshal(message, &event)

	return event, err
}

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
