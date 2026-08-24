package math_tool

import (
	"github.com/odin0528/tmg-common/math_rng"
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

func PickByWeightsByPool(weights []float64, oddLimit float64) (pickIdx int) {
	return math_rng.PickByWeights(weights)
}
