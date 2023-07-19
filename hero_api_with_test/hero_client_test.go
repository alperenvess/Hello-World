package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeroClint(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, req.URL.String(), "/")
			rw.Write([]byte(`{
				"squadName": "Xyz"
				"townName": "Warsaw"

			}`))
		}),
	)

	defer server.Close()

	api := HeroClient{Client: server.Client(),
		Url: server.URL}

	r, err := api.GetThem()

	assert.Nil(t, err)

	assert.Equal(t, r.SquadName, "Xyz")
	assert.Equal(t, r.HomeTown, "Warsaw")

	fmt.Printf("%v", r)
}
