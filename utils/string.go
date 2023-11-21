package utils

import "game_server/common/math_tool"

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
