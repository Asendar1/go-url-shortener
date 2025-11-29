package handler

import (
	"net/http"

	store "github.com/Asendar1/go-url-shortener/store"
	utils "github.com/Asendar1/go-url-shortener/utils"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "This URI only for POST method", http.StatusMethodNotAllowed)
		return
	}
	// ? Testing for now
	origin_url := r.FormValue("url")
	if origin_url == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}
	url_pair := utils.FormURL(origin_url)
	store.SaveURLPair(url_pair)
}
