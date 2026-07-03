package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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

func GenerateSignature(token, method, path, timestamp string) string {
	payload := fmt.Sprintf("%s|%s|%s", method, path, timestamp)
	h := hmac.New(sha256.New, []byte(token))
	h.Write([]byte(payload))

	return hex.EncodeToString(h.Sum(nil))
}
