package pokeapi

import (
	"fmt"
	"net/url"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int64  `json:"base_experience"`
}

func (c *Client) GetPokemon(name *string) (Pokemon, error) {
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

func (c *Client) CatchPokemon(name *string) (Pokemon, error) {
	pokemon, err := c.GetPokemon(name)
	if err != nil {
		return pokemon, fmt.Errorf("error catching pokemon: %w", err)
	}
	printPokemonCaught(pokemon)
	return pokemon, nil
}

func printPokemonCaught(pokemon Pokemon) {
	fmt.Println("pokemon was caught!")
	fmt.Println(pokemon)
}
