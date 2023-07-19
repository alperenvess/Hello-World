package main

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	ServiceName string
}

func (app *App) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World from "+app.ServiceName)
}

func main() {
	app := App{ServiceName: "MyApp"}
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.HelloWorld)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
