package utils

import (
	"encoding/json"
	"errors"
	"log"
	math_rand "math/rand"
	"time"
	"xxx/common/math_tool"
	"xxx/common/uid"
	"xxx/common/web/response"

	"github.com/seehuhn/mt19937"
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

func PickByWeights(weights []float64) (idx int) {
	if len(weights) == 0 {
		log.Println("weights is empty")
		return 0
	}

	src := math_rand.New(mt19937.New())
	src.Seed(int64(uid.GenerateUniqueID()) + time.Now().UnixNano())
	weightHandler := math_tool.NewWeighted(weights, src)

	idx, _ = weightHandler.Take()

	return idx
}
