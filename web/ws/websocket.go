package ws

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"xxx/common/configs"
	"xxx/common/math_tool"

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
	} else if encodeType == configs.ENABLE_WS_MIX_BASE64_SHIFT {
		encode := EncodeByBase64Shift(byteArray)
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
	} else if decodeType == configs.ENABLE_WS_MIX_BASE64_SHIFT {
		decode, _ := DecodeByBase64Shift(message)
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

func EncodeByBase64Shift(src []byte) []byte {
	// base64 > shift > add randNum(2位組的16進制)在字段最前方 > base64
	encode := EncodeBybase64(src)

	randomNum := 0
	if len(encode) > HEX_MAX_BIT {
		randomNum = math_tool.GetRandInt(HEX_MAX_BIT) + 1
	} else {
		randomNum = math_tool.GetRandInt(len(encode)) + 1
	}

	shiftLength := randomNum

	if shiftLength < len(encode) {
		shiftedResult := append(encode[len(encode)-shiftLength:], encode[:len(encode)-shiftLength]...)
		encode = shiftedResult
	}

	hexStr := fmt.Sprintf("%02x", shiftLength) // 補0 確保為2位組

	encode = append([]byte(hexStr), encode...)

	return EncodeBybase64(encode)
}

func DecodeByBase64Shift(src []byte) ([]byte, error) {
	var data []byte

	if len(src) == 0 {
		return data, errors.New("message is empty")
	}

	decode, _ := DecodeByBase64(src)

	randomNum := string(decode)[:2]

	shiftLength64, err := strconv.ParseInt(randomNum, 16, 64)
	if err != nil {
		return data, err
	}
	shiftLength := int(shiftLength64)

	decode = []byte(string(decode)[2:])

	if shiftLength < len(decode) {
		shiftedResult := append(decode[shiftLength:], decode[:shiftLength]...)
		decode = shiftedResult
	}

	return DecodeByBase64(decode)
}
