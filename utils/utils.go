package utils

import (
	"log"
	"strconv"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func IsAlphanumeric(s string) bool {
	return alphanumericRegex.MatchString(s)
}

func BoolToInt(boolean bool) (integer int) {
	if boolean {
		integer = 1
	}
	return integer
}

func SliceBoolToInt(booleans []bool) (integers []int) {
	for _, v := range booleans {
		integers = append(integers, BoolToInt(v))
	}
	return integers
}

func IsElementExist[T comparable](list []T, element T) bool {
	for _, e := range list {
		if e == element {
			return true
		}
	}
	return false
}

func StringToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

func TernaryOperatorString(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}

func TernaryOperatorInt(condition bool, trueVal, falseVal int) int {
	if condition {
		return trueVal
	}
	return falseVal
}
