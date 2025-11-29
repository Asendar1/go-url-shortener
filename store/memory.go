package store

import models "github.com/Asendar1/go-url-shortener/models"

var (
	ShortToLong = make(map[string]string)
	LongToShort = make(map[string]string)
)

func SaveURLPair(pair models.URLPair) {
	ShortToLong[pair.ShortURL] = pair.OriginalURL
	LongToShort[pair.OriginalURL] = pair.ShortURL // optional
}

func FindOriginalURL(shortURL string) (string, bool) {
	rtn, err := ShortToLong[shortURL]
	return rtn, err
}
