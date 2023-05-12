package utils

import "xxx/common/math_tool"

func GetRandomString(length int) string {
	b := make([]rune, length)
	for idx := range b {
		b[idx] = letterRunes[math_tool.GetRandomInt(len(letterRunes))]
	}
	return string(b)
}
