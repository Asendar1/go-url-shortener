package utils

import (
	"math/rand"
	"time"

	models "github.com/Asendar1/go-url-shortener/models"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var src = rand.NewSource(time.Now().UnixNano())


func FormURL(url string) models.URLPair {
	r := rand.New(src)
	b := make([]byte, 6)

	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}

	return models.URLPair{
		OriginalURL: url,
		ShortURL   : string(b),
	}
}
