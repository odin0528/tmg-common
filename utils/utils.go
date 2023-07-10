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
