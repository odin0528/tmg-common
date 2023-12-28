package utils

import "XXX/common/math_tool"

func GetRandomString(length int) string {
	b := make([]rune, length)
	for idx := range b {
		b[idx] = letterRunes[math_tool.GetRandomInt(len(letterRunes))]
	}
	return string(b)
}

func IsExist(source []string, target string) bool {
	for i := 0; i < len(source); i++ {
		if source[i] == target {
			return true
		}
	}

	return false
}

func Remove(source []string, target string) []string {
	for i, v := range source {
		if v == target {
			return append(source[:i], source[i+1:]...)
		}
	}
	return source
}
