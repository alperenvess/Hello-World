package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
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

			"en": "hello welcome",

			"pl": "dzień dobry",

			"de": "guten tag",
		}

		if msg, ok := helloMsgs[lang]; ok {

			w.Write([]byte(msg + " " + name))

			fmt.Fprintf(w, msg+" from "+helloMsgs[lang]+" to "+name)

		} else {

			w.WriteHeader(404)

			w.Write([]byte("Unknown language"))

		}

	})

	func TestByHandlers(t *testing.T) {

		r, err := http.NewRequest("GET", "http://test.com?lang=pl&user=alp", nil)

		w := httptest.NewRecorder()

		if w.Code != http.StatusOK {
			t.Fatal(w.Code, w.Body.String())
		}
	}

	http.ListenAndServe(":3000", r)

}
