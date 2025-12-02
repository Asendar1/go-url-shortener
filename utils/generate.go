package utils

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var src = rand.NewSource(time.Now().UnixNano())


func FormURL(url string) string {
	r := rand.New(src)
	b := make([]byte, 6)

	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}

	return string(b)
}
