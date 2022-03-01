package service

import (
	"fmt"
	"testing"
)

func TestWeatherbyCityClient_WeatherbyCity(t *testing.T) {

	tests := []struct {
		name string
	}{
		{name: "Wrong input"},
		{name: "Right input"},
	}

	for _, tt := range tests {
		fmt.Println(tt)
	}

}
