package pokeapi

import (
	"fmt"
	"net/url"
)

type LocationAreaExplored struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ExploreLocation(locationArea *string) (LocationAreaExplored, error) {
	reqUrl, _ := url.JoinPath(BaseUrl, LocationEndpoint, *locationArea)
	response := LocationAreaExplored{}
	data, err := c.GetUrl(reqUrl)
	if err != nil {
		return response, fmt.Errorf("error getting location area: %w", err)
	}
	UnmarshalData(data, &response)
	return response, nil
}
