package math_tool

import (
	"math/rand"
	"time"
)

func SetUpRandSeed() {
	randomOnce.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})
}

func GetRandomFloat64(max float64) float64 {
	return float64(GetRandomInt64(int64(max)))
}

func GetRandomInt(max int) int {
	return int(GetRandomInt64(int64(max)))
}

func GetRandomInt64(max int64) int64 {
	if max <= 0 {
		return 0
	}
	return rand.Int63n(max)
}

func Shuffle[T any](target []T) {
	rand.Shuffle(len(target), func(i, j int) {
		target[i], target[j] = target[j], target[i]
	})
}

func GetShuffleCopyArray[T any](target []T) []T {
	n := make([]T, len(target))
	copy(n, target)
	Shuffle(n)
	return n
}
