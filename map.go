package main

import (
	"fmt"

	"github.com/aramirez3/pokedexcli/internal/pokeapi"
)

func commandMapF(cfg *config) error {
	var currentUrl string
	if cfg.Next == nil {
		currentUrl = pokeapi.BaseUrl
	} else {
		currentUrl = *cfg.Next

	}

	locationsResponse := pokeapi.LocationsResponse{}
	cached, found := cfg.Cache.Get(currentUrl)
	if found {
		pokeapi.UnmarshalData(cached, &locationsResponse)
		cfg.Next = locationsResponse.Next
		cfg.Previous = locationsResponse.Previous
		prinLocationAreas(locationsResponse.Results)
		return nil
	}

	locationAreasResponse, err := cfg.pokeapiClient.GetLocations(cfg.Next)
	if err != nil {
		return fmt.Errorf("error getting location-areas: %w", err)
	}

	byteData, _ := pokeapi.MarshalData(locationAreasResponse)
	cfg.Cache.Add(currentUrl, byteData)
	cfg.Next = locationAreasResponse.Next
	cfg.Previous = locationAreasResponse.Previous
	prinLocationAreas(locationAreasResponse.Results)
	return nil
}

func commandMapB(cfg *config) error {
	var currentUrl string
	if cfg.Previous == nil {
		currentUrl = pokeapi.BaseUrl
	} else {
		currentUrl = *cfg.Previous
	}

	locationsResponse := pokeapi.LocationsResponse{}
	cached, found := cfg.Cache.Get(currentUrl)

	if found {
		pokeapi.UnmarshalData(cached, &locationsResponse)
		cfg.Next = locationsResponse.Next
		cfg.Previous = locationsResponse.Previous
		prinLocationAreas(locationsResponse.Results)
		return nil
	}

	locationAreas, err := cfg.pokeapiClient.GetLocations(&currentUrl)
	if err != nil {
		return fmt.Errorf("error getting location-areas: %w", err)
	}

	byteData, _ := pokeapi.MarshalData(locationAreas.Results)
	cfg.Cache.Add(currentUrl, byteData)
	cfg.Next = locationAreas.Next
	cfg.Previous = locationAreas.Previous
	prinLocationAreas(locationAreas.Results)
	return nil
}

func prinLocationAreas(locationAreas []pokeapi.LocationArea) {
	for _, loc := range locationAreas {
		fmt.Println(loc.Name)
	}
}
