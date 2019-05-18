package utils

import (
	"math/rand"
	"time"
)

//nolint:gochecknoinits
func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	//nolint:gochecknoglobals
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(letterRunes[rand.Intn(len(letterRunes))])
	}
	return b
}
