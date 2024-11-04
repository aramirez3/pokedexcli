package pokeapi

import (
	"fmt"
)

type LocationsResponse struct{}

func getLocations(param string) (LocationsResponse, error) {
	locationsResponse := LocationsResponse{}
	url := baseUrl
	if param != "" {
		url = baseUrl + param
	}

	fmt.Printf("url: %s", url)
	return locationsResponse, nil
}
