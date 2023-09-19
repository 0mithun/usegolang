package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my great site updat<h1>")
}

func main() {
	// http.HandleFunc("/", handlerFunc)

	r := chi.NewRouter()
	r.Get("/", homeHandler)

	fmt.Println("Starting server on port 3000")
	http.ListenAndServe("localhost:3000", r)
}
