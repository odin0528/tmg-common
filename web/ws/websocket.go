package ws

import (
	"encoding/base64"
	"encoding/json"
	"mgmt/common/configs"

	"github.com/vmihailenco/msgpack/v5"
)

func GetEventResponse(event string, code int, msg string, data interface{}) []byte {
	rsp := EventResponse{
		Event: event,
		Msg:   msg,
		Code:  code,
		Data:  data,
	}

	byteArray, _ := json.Marshal(rsp)

	if encodeType := configs.GetInt(configs.SECTION_SYSTEM, configs.SYSTEM_WEBSOCKET_ENCODE_MODE, configs.ENABLE_WS_BASE64); encodeType == configs.ENABLE_WS_BASE64 {
		encode := EncodeBybase64(byteArray)
		return encode
	} else if encodeType == configs.ENABLE_WS_MSG_PACK {
		encode, _ := EncodeByMsgpack(byteArray)
		return encode
	}

	return byteArray
}

func GetServerErrorResponse(code int, msg string) []byte {
	return GetEventResponse(EVENT_WS_SERVER_ERROR, code, msg, nil)
}

func ParseEvent(message []byte) (Event, error) {
	var event Event

	if decodeType := configs.GetInt(configs.SECTION_SYSTEM, configs.SYSTEM_WEBSOCKET_DECODE_MODE, configs.ENABLE_WS_BASE64); decodeType == configs.ENABLE_WS_BASE64 {
		decode, _ := DecodeByBase64(message)
		err := json.Unmarshal(decode, &event)
		return event, err
	} else if decodeType == configs.ENABLE_WS_MSG_PACK {
		decode, _ := DecodeByMsgpack(message)
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

func EncodeByMsgpack(src []byte) ([]byte, error) {
	return msgpack.Marshal(src)
}

func DecodeByMsgpack(src []byte) ([]byte, error) {
	var data []byte
	err := msgpack.Unmarshal(src, &data)
	return data, err
}
