package main

import (
	"log"
	"net/http"

	handler "github.com/Asendar1/go-url-shortener/handlers"
	"github.com/Asendar1/go-url-shortener/store"
)


func main() {
	mux := http.NewServeMux()
	URLStore, err := store.Connect("postgres://dev:dev@localhost:5432/urlshortener?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to the db:", err)
	}
	handler.SetStore(URLStore)
	// POST and PUT and DELETE Handlers
	mux.HandleFunc("/shorten", handler.Shorten)
	mux.HandleFunc("/shorten/", handler.Shorten)
	// GET Handlers
	mux.HandleFunc("/", handler.Redirect)
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
