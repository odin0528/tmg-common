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

func StringIntersection(slice1, slice2 []string) []string {
	set := make(map[string]bool)
	result := []string{}

	for _, item := range slice1 {
		set[item] = true
	}

	for _, item := range slice2 {
		if set[item] {
			result = append(result, item)
			delete(set, item)
		}
	}

	return result
}

func Remove(source []string, target string) []string {
	for i, v := range source {
		if v == target {
			return append(source[:i], source[i+1:]...)
		}
	}
	return source
}

func GetCombinedString(symbol string, keys ...string) string {
	result := ""
	for i, key := range keys {
		result += key

		if i != (len(keys) - 1) {
			result += symbol
		}
	}

	return result
}
