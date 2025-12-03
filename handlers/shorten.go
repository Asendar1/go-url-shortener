package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	store "github.com/Asendar1/go-url-shortener/store"
	utils "github.com/Asendar1/go-url-shortener/utils"
)

var URLStore *store.Store

func SetStore(store *store.Store) {
	URLStore = store
}

// Helper function to handle URL shortening

func HandleGETStats(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[len("/shorten/stats/"):]
	shortCode = strings.TrimSpace(shortCode)
	if shortCode == "" {
		utils.JSONError(w, http.StatusBadRequest, "No short code provided")
		return
	}
	urlDb, err := URLStore.GetByShortCode(shortCode)
	if err != nil {
		utils.JSONError(w, http.StatusNotFound, "Not Found in Database")
		return
	}
	utils.JSONSuccess(w, 200, map[string]any{
		"id": 		fmt.Sprintf("%d", urlDb.ID),
		"url":		urlDb.LongUrl,
		"short_code":	urlDb.ShortCode,
		"created_at":	urlDb.CreatedAt.Time.String(),
		"clicks":	urlDb.Clicks.Int64,
	})
}

func HandleCreateShortURL(w http.ResponseWriter, r *http.Request) {
	origin_url := r.FormValue("url")
	if origin_url == "" {
		utils.JSONError(w, http.StatusBadRequest, "No URL has been given")
		return
	}
	var tries int;
	short_url := utils.FormURL(origin_url)
	for {
		found := URLStore.CreateURL(short_url, origin_url); if found == nil {
			break
		}
		short_url = utils.FormURL(origin_url)
		if tries > 10 {
			utils.JSONError(w, http.StatusInternalServerError, found.Error())
			return
		}
		tries++
	}
	url_db, _ := URLStore.GetByShortCode(short_url)
	utils.JSONSuccess(w, http.StatusCreated, map[string]string{
		"id":		fmt.Sprintf("%d", url_db.ID),
		"url":		url_db.LongUrl,
		"short_code":	url_db.ShortCode,
		"created_at":	url_db.CreatedAt.Time.String(),
	},
	)
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet{
		HandleGETStats(w, r)
		return
	}
	if r.Method != http.MethodPost && r.Method != http.MethodPut && r.Method != http.MethodDelete {
		utils.JSONError(w, http.StatusBadRequest, "Bad Request Method")
		return
	}
	if r.Method == http.MethodPut {
		UpdateShortUrl(w, r)
		return
	}
	if r.Method == http.MethodDelete {
		DeleteURL(w, r)
		return
	}
	HandleCreateShortURL(w, r)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.JSONError(w, http.StatusBadRequest, "Bad Request Method")
		return
	}
	short_url := r.URL.Path[1:]
	url_db, err := URLStore.GetByShortCode(short_url)
	URLStore.UpdateClicks(short_url)
	if err == nil {
		utils.JSONSuccess(w, 200, map[string]string {
			"id": 	fmt.Sprintf("%d", url_db.ID),
			"url": url_db.LongUrl,
			"short_code": short_url,
			"created_at": url_db.CreatedAt.Time.String(),
		})
		return
	}
	utils.JSONError(w, http.StatusNotFound, "Not Found in Database")
}

func UpdateShortUrl(w http.ResponseWriter, r *http.Request) {
	var urlToUpdate struct {
		URL string	`json:"url"`
	}
	json.NewDecoder(r.Body).Decode(&urlToUpdate)
	if urlToUpdate.URL == ""{
		utils.JSONError(w, http.StatusBadRequest, "Both short URL and new long URL must be provided")
		return
	}
	short_code := r.URL.Path[len("/shorten/"):]
	err := URLStore.UpdateLongUrl(short_code, urlToUpdate.URL)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Failed to update URL")
		return
	}
	url_db, _ := URLStore.GetByShortCode(short_code)
	utils.JSONSuccess(w, 200, map[string]string {
			"id": 	fmt.Sprintf("%d", url_db.ID),
			"url": url_db.LongUrl,
			"short_code": short_code,
			"created_at": url_db.CreatedAt.Time.String(),
		})
}

func DeleteURL(w http.ResponseWriter, r *http.Request) {
	short_code := r.URL.Path[len("/shorten/"):]
	err := URLStore.DeleteByShortCode(short_code)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Failed to delete URL")
		return
	}
	utils.JSONSuccess(w, 200, map[string]string{
		"message": "URL deleted successfully",
	})
}
