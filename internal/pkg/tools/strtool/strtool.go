package strtool

import (
	"math/rand"
	"strings"
	"time"
)

func TrimRightSpace(s string) string {
	return strings.TrimRight(string(s), "\r\n\t ")
}

// Create a random string with a given length
func RandomString(length int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// CompareStrings compares two strings and returns true if they are equal,
// otherwise returns false. It performs a case-sensitive comparison by default.
func CompareStrings(str1, str2 string) bool {
	return str1 == str2
}

// CompareStringsIgnoreCase compares two strings case-insensitively and returns true if they are equal,
// otherwise returns false.
func CompareStringsIgnoreCase(str1, str2 string) bool {
	return strings.EqualFold(str1, str2)
}
