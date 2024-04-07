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
