package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {

		name := r.URL.Query().Get("name")

		lang := r.URL.Query().Get("lang")

		helloMsgs := map[string]string{

			"en": "hello",

			"pl": "cześć",
		}

		if msg, ok := helloMsgs[lang]; ok {

			w.Write([]byte(msg + " " + name))

		} else {

			w.WriteHeader(404)

			w.Write([]byte("Unknown language"))

		}

	})

	http.ListenAndServe(":3000", r)

}
