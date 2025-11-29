package main

import (
	"fmt"
	"net/http"

	handler "github.com/Asendar1/go-url-shortener/handlers"
)

func main() {
	// TODO : remeber to change this before deployment
	fmt.Println("Starting URL test server on localhost:8080")
	mux := http.NewServeMux()

	// POST Handlers
	mux.HandleFunc("/shorten", handler.Shorten)
	mux.HandleFunc("/shorten/", handler.Shorten)
	// GET Handlers
	mux.HandleFunc("/", handler.Redirect)

	http.ListenAndServe("localhost:8080", mux)
}
