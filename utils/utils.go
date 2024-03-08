package utils

import "log"

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
