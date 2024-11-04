package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationsResponse struct{}

func (c *Client) GetLocations(param string) (LocationsResponse, error) {
	locationsResponse := LocationsResponse{}
	url := baseUrl
	if param != "" {
		url = baseUrl + param
	}

	fmt.Printf("url: %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationsResponse, fmt.Errorf("error making request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationsResponse, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationsResponse, fmt.Errorf("error encoding data: %w", err)
	}

	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return locationsResponse, fmt.Errorf("error unmarshaling data: %w", err)
	}

	return locationsResponse, nil
}
