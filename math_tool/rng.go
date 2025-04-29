package math_tool

import (
	"gltw.6633663.com/Backend_GSGame/common/math_rng"
)

func GetRandomInt64(max int64) int64 {
	return math_rng.GetRandomInt64(max)
}

func GetRandomInt(max int) int {
	return math_rng.GetRandomInt(max)
}

func Shuffle[T any](target []T) {
	math_rng.Shuffle(target)
}

func PickByWeights(weights []float64) (pickIdx int) {
	return math_rng.PickByWeights(weights)
}
