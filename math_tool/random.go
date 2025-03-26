package math_tool

import (
	math_rand "math/rand"
)

func GetRandomFloat64(max float64) float64 {
	return float64(GetRandomInt64(int64(max)))
}

func GetShuffleCopyArray[T any](target []T) []T {
	n := make([]T, len(target))
	copy(n, target)
	Shuffle(n)
	return n
}

// 取區間: min ~ max-1 隨機數(不包含max)
func GetRandomRangeNum(r *math_rand.Rand, min, max int) int {
	var result int
	switch {
	case min > max:
		min, max = max, min
		maxRand := max - min
		b := r.Intn(maxRand)
		result = min + b
	case max == min:
		result = max
	case max > min:
		maxRand := max - min
		b := r.Intn(maxRand)
		result = min + b
	}

	return result
}

// 取區間: min ~ max 隨機數(包含max)
func GetRandomRangeNumWithMax(r *math_rand.Rand, min, max int) int {
	return GetRandomRangeNum(r, min, max+1)
}
