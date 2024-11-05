package main

import (
	"fmt"

	"github.com/aramirez3/pokedexcli/internal/pokeapi"
)

func commandMapF(cfg *config) error {
	locationAreas := []pokeapi.LocationArea{}
	cached, found := cfg.Cache.Get(*cfg.Next)
	if found {
		pokeapi.UnmarshalData(cached, &locationAreas)
	} else {
		locationAreasResponse, err := cfg.pokeapiClient.GetLocations(cfg.Next)
		if err != nil {
			return fmt.Errorf("error getting location-areas: %w", err)
		}
		cfg.Next = locationAreasResponse.Next
		cfg.Previous = locationAreasResponse.Previous
		locationAreas = locationAreasResponse.Results
	}
	for _, loc := range locationAreas {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	locationAreas, err := cfg.pokeapiClient.GetLocations(cfg.Previous)
	if err != nil {
		return fmt.Errorf("error getting location-areas: %w", err)
	}
	cfg.Next = locationAreas.Next
	cfg.Previous = locationAreas.Previous
	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
