package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	c := &http.Client{
		Timeout: 2 * time.Second,
	}

	log.Println("Fetching...")
	resp, err := c.Get(
		"https://mdn.github.io/learning-area/javascript/oojs/json/superheroes.json")

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
