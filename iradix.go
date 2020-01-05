package benchtrie

import (
	"fmt"
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func RandString(length int) string {
	return StringWithCharset(length, charset)
}

func GenerateHosts(size int) (result []string) {
	for i := 0; i < size; i++ {
		result = append(result, fmt.Sprintf("ads.%d.%s.com", i, RandString(5)))
	}
	return
}
