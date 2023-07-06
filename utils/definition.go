package utils

import (
	"regexp"
)

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var alphanumericRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)