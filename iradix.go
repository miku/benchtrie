package benchtrie

import (
	"fmt"
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyz"
const numset = "0123456789"

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

// GenerateHosts returns a number of random hostnames. The TLD might be short,
// so to increase duplicate TLDs, which is more common.
func GenerateHosts(size int) (result []string) {
	for i := 0; i < size; i++ {
		result = append(result, fmt.Sprintf("ads.%d.%s.com", i, RandString(3)))
	}
	return
}

func GenerateIPv4Addr(size int) (result []string) {
    for i := 0; i < size; i++ {
        result = append(result, fmt.Sprintf("%s.%s.%s.%s", RandString(3),
                        RandString(3), RandString(3), RandString(3)))
    }
    return
}
