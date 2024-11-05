package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type LocationsResponse struct {
	Count    int64
	Next     *string
	Previous *string
	Results  []LocationArea
}

type LocationArea struct {
	Name string
	Url  string
}

func (c *Client) GetLocations(param *string) (LocationsResponse, error) {
	locationsResponse := LocationsResponse{}
	reqUrl, _ := url.JoinPath(baseUrl, locationArea)
	if param != nil {
		reqUrl = *param
	}

	req, err := http.NewRequest("GET", reqUrl, nil)
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

	UnmarshalData(data, &locationsResponse)
	// err = json.Unmarshal(data, &locationsResponse)
	// if err != nil {
	// 	return locationsResponse, fmt.Errorf("error unmarshaling data: %w", err)
	// }

	return locationsResponse, nil
}
