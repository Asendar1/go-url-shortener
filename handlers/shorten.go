package handler

import (
	"net/http"

	store "github.com/Asendar1/go-url-shortener/store"
	utils "github.com/Asendar1/go-url-shortener/utils"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.JSONError(w, http.StatusBadRequest, "Only POST method is allowed")
		return
	}
	origin_url := r.FormValue("url")
	if origin_url == "" {
		utils.JSONError(w, http.StatusBadRequest, "No URL has been given")
		return
	}
	if _, found := store.LongToShort[origin_url]; found {
		utils.JSONError(w, http.StatusConflict, "This URL already been")
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
			utils.JSONError(w, http.StatusInternalServerError, "Could not generate please try again")
		}
		tries++
	}
	store.SaveURLPair(url_pair)
	// return body
	utils.JSONSuccess(w, http.StatusCreated, map[string]string {
		"url": url_pair.OriginalURL,
		"short_code": url_pair.ShortURL,
		},
	)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	short_url := r.URL.Path[1:]
	original_url, found := store.FindOriginalURL(short_url)
	if found {
		utils.JSONSuccess(w, 200, map[string]string {
			"url": original_url,
			"short_code": short_url,
		})
		return
	}
	utils.JSONError(w, http.StatusNotFound, "Not Found in Database")
}
