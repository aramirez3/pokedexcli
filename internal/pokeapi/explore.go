package pokeapi

import (
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
	reqUrl, _ := url.JoinPath(BaseUrl, Location, *locationArea)
	response := LocationAreaExplored{}
	data, _ := c.GetUrl(reqUrl)
	UnmarshalData(data, &response)
	return response, nil
}
