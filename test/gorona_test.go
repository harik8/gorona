package main

import (
	"testing"

	"github.com/harik8/gorona/client"
)

var countries = []string{"China", "Italy", "USA", "Netherlands", "Sri Lanka"}

func TestGetCountries(t *testing.T) {
	client.GetCountries()
}
func TestGetCountry(t *testing.T) {
	for i := 0; i < len(countries); i++ {
		client.GetCountry(countries[i])
	}
}
