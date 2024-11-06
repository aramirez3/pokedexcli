package pokeapi

import (
	"fmt"
	"net/url"
)

type Pokemon struct {
	name           string `json:"name"`
	baseExperience int64  `json:"base_experience"`
}

func (c *Client) getPokemon(name *string) (Pokemon, error) {
	pokemonResponse := Pokemon{}
	reqUrl, err := url.JoinPath(BaseUrl, PokemonEndpoint)

	if name != nil {
		reqUrl, err = url.JoinPath(BaseUrl, PokemonEndpoint, *name)
	}

	if err != nil {
		return pokemonResponse, fmt.Errorf("error joining path: %w", err)
	}

	data, err := c.GetUrl(reqUrl)

	if err != nil {
		return pokemonResponse, fmt.Errorf("error getting pokemon: %w", err)
	}
	UnmarshalData(data, &pokemonResponse)
	return pokemonResponse, nil
}
