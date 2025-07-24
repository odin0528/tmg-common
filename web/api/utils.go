package api

import "strings"

func GetUrl(address string, path string) string {
	return address + path
}

func IsValidResponseBody(dataByte []byte) bool {
	if len(dataByte) == 0 || strings.Compare(string(dataByte), "{}") == 0 || strings.Compare(string(dataByte), "null") == 0 {
		return false
	}

	return true
}
