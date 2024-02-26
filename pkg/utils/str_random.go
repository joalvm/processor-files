package utils

import (
	"math/rand"
	"time"
)

func StrRandom(length int) string {
	const charset = "abcdef0123456789"

	if length == 0 {
		length = 16
	}

	b := make([]byte, length)
	rand.NewSource(time.Now().UnixNano())
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
