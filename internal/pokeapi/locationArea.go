package pokeapi

import (
	"fmt"
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
	reqUrl, err := url.JoinPath(BaseUrl, LocationEndpoint)
	if err != nil {
		return locationsResponse, fmt.Errorf("error getting location: %w", err)
	}
	if param != nil {
		reqUrl = *param
	}
	data, _ := c.GetUrl(reqUrl)

	UnmarshalData(data, &locationsResponse)

	return locationsResponse, nil
}
