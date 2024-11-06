package pokeapi

import (
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
	reqUrl, _ := url.JoinPath(BaseUrl, Location)
	if param != nil {
		reqUrl = *param
	}
	locationsResponse := LocationsResponse{}
	data, _ := c.GetUrl(reqUrl)

	UnmarshalData(data, &locationsResponse)

	return locationsResponse, nil
}
