package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello!")
	})

	mux.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "World!")
	})

	err := http.ListenAndServe(":3000", mux)

	log.Fatal(err)
}
