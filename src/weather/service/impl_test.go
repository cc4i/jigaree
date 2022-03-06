package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	server *httptest.Server
)

func TestWeatherbyCityClient_WeatherbyCity(t *testing.T) {

	tests := []struct {
		name string
	}{
		{name: "Wrong input"},
		{name: "Beijing input"},
	}

	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// mock here
	}))

	fmt.Println(server.URL)
	for _, tt := range tests {
		fmt.Println(tt)
	}

}
