package utils

import (
	"errors"
	"fmt"
	"mgmt/common/math_tool"
	"strconv"
)

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

func ConvertHashToIntegerList(hash string, charsPerGroup int) ([]int, error) {
	b := []byte(hash)
	if len(b) != 64 {
		return nil, errors.New("invalid hash length")
	}

	if len(b)%charsPerGroup != 0 {
		return []int{}, fmt.Errorf("hash length (%d) is not divisible by charsPerGroup (%d)", len(b), charsPerGroup)
	}

	integerList := make([]int, 0, len(b)/charsPerGroup)
	for i := 0; i < len(b); i += charsPerGroup {
		integer, err := strconv.ParseInt(string(b[i:i+charsPerGroup]), 16, 64)
		if err != nil {
			return []int{}, err
		}

		integerList = append(integerList, int(integer))
	}

	return integerList, nil
}
