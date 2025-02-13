package utils

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"game_server/common/math_tool"
	"game_server/common/web/response"
	"log"

	"math"
	"math/big"
)

func ToGenericSlice[T any](input []T) []any {
	result := make([]any, len(input))
	for i, v := range input {
		result[i] = v
	}
	return result
}

func ToTypedSlice[T any](input []any) ([]T, error) {
	result := make([]T, len(input))
	for i, v := range input {
		if val, ok := v.(T); ok {
			result[i] = val
		} else {
			return nil, fmt.Errorf("invalid type conversion at index %d", i)
		}
	}
	return result, nil
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

func PickByWeights(weights []float64) (pickIdx int) {
	if len(weights) == 0 {
		log.Println("weights is empty")
		return 0
	}

	sum := 0.0
	maxLimitList := make([]float64, len(weights))

	for idx, weight := range weights {
		sum += weight
		maxLimitList[idx] = sum
	}

	sum *= math.Pow10(6)

	bigInt := new(big.Int).SetInt64(int64(sum))
	v, _ := rand.Int(rand.Reader, bigInt)

	value := float64(v.Int64()) * math.Pow10(-6)

	pickIdx = 0
	for idx, limit := range maxLimitList {
		if math_tool.IsFloatLessThan(value, limit) {
			pickIdx = idx
			break
		}
	}

	return pickIdx
}
