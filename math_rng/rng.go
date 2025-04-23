package math_rng

import (
	"crypto/rand"
	"log"
	"math"
	"math/big"
)

func GetRandomInt64(max int64) int64 {
	if max <= 0 {
		return 0
	}
	bigInt := new(big.Int).SetInt64(int64(max))
	i, _ := rand.Int(rand.Reader, bigInt)
	return i.Int64()
}

func GetRandomInt(max int) int {
	return int(GetRandomInt64(int64(max)))
}

func Shuffle[T any](target []T) {
	for i := range target {
		j := GetRandomInt64(int64(i) + 1)
		target[i], target[j] = target[j], target[i]
	}
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
		if value < limit {
			pickIdx = idx
			break
		}
	}

	return pickIdx
}
