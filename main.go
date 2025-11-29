package main

import (
	"log"
	"net/http"

	handler "github.com/Asendar1/go-url-shortener/handlers"
)

func main() {
	mux := http.NewServeMux()

	// POST Handlers
	mux.HandleFunc("/shorten", handler.Shorten)
	mux.HandleFunc("/shorten/", handler.Shorten)
	// GET Handlers
	mux.HandleFunc("/", handler.Redirect)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
