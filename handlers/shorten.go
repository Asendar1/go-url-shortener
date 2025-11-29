package handler

import (
	"fmt"
	"net/http"

	store "github.com/Asendar1/go-url-shortener/store"
	utils "github.com/Asendar1/go-url-shortener/utils"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "This URI only for POST method", http.StatusMethodNotAllowed)
		return
	}
	origin_url := r.FormValue("url")
	if origin_url == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}
	if _, found := store.LongToShort[origin_url]; found {
		http.Error(w, "URL has already been shortened", http.StatusConflict)
		return
	}
	var tries int;
	url_pair := utils.FormURL(origin_url)
	for {
		if _ , found := store.ShortToLong[url_pair.ShortURL]; !found {
			break
		}
		url_pair = utils.FormURL(origin_url)
		if tries > 10 {
			http.Error(w, "Could not regerenate short url, please try again", http.StatusInternalServerError)
		}
		tries++
	}
	store.SaveURLPair(url_pair)
	fmt.Fprintf(w, "%s", url_pair.ShortURL)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	short_url := r.URL.Path[1:]
	fmt.Printf("short before: %s\n", short_url)
	original_url, found := store.FindOriginalURL(short_url)
	if found {
		http.Redirect(w, r, original_url, http.StatusPermanentRedirect)
		return
	}
	fmt.Printf("short:%s long:%s %v", short_url, original_url, found)
	http.Error(w, "invalid URL", http.StatusBadRequest)
}
