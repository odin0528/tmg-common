package utils

import (
	"crypto/sha256"
	"fmt"
)

func GetSHA256Hash(input string) string {
	sum := sha256.Sum256([]byte(input))
	result := fmt.Sprintf("%x", sum)
	return result
}

func CheckSHA256Hash(input, hash string) bool {
	inputHash := GetSHA256Hash(input)
	return inputHash == hash
}

func GetAPIKeyHash(platformName, privateKey, timeStamp string) string {
	temp := platformName + "-" + privateKey + "-" + timeStamp
	return GetSHA256Hash(temp)
}