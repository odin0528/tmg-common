package math_tool

import (
	"crypto/rand"
	"math/big"
)

func GetRandInt(max int) int {
	return int(GetRandInt64(max))
}

func GetRandInt64(max int) int64 {
	bigInt := new(big.Int).SetInt64(int64(max))
	i, _ := rand.Int(rand.Reader, bigInt)
	return i.Int64()
}

func GetShuffleArray[T any](target []T) {

	for i := range target {
		j := GetRandInt64(i + 1)
		target[i], target[j] = target[j], target[i]
	}

}
