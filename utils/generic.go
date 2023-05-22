package utils

import (
	"encoding/json"
	"errors"
	"game_server/common/web/response"
)

func ToGenericSlice[T any](input []T) []any {
	result := make([]any, len(input))
	for i, v := range input {
		result[i] = v
	}
	return result
}

func ToStringSpecifiedTypeMap[T any](input map[string]interface{}) (map[string]T, error) {
	result := map[string]T{}
	for i, v := range input {

		jsonByte, err := json.Marshal(v)
		if err != nil {
			return nil, errors.New(response.MSG_MARSHAL_ERROR)
		}

		var tv T
		err = json.Unmarshal(jsonByte, &tv)
		if err != nil {
			return nil, errors.New(response.MSG_UNMARSHAL_ERROR)
		}
		result[i] = tv
	}

	return result, nil
}
