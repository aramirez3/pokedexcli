package main

import (
	"fmt"

	"github.com/aramirez3/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config) error {
	areaName := ""
	if len(cfg.Words) > 0 {
		areaName = cfg.Words[0]
	} else {
		fmt.Println("explore requires an area-name")
		return nil
	}
	cached, found := cfg.Cache.Get(areaName)
	if found {
		locationExplored := pokeapi.LocationAreaExplored{}
		pokeapi.UnmarshalData(cached, &locationExplored)
		printLocationExplored(locationExplored)
		return nil
	}

	locationExplored, err := cfg.pokeapiClient.ExploreLocation(&areaName)
	if err != nil {
		return fmt.Errorf("error exploring area: %w", err)
	}

	byteData, err := pokeapi.MarshalData(locationExplored)
	if err != nil {
		return fmt.Errorf("error marshaling explore data: %w", err)
	}
	cfg.Cache.Add(areaName, byteData)

	addEncountersToCurrentRange(locationExplored, cfg)

	printLocationExplored(locationExplored)
	return nil
}

func printLocationExplored(locationExplored pokeapi.LocationAreaExplored) {
	fmt.Printf("Exploring %s...\n", locationExplored.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationExplored.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
}

func addEncountersToCurrentRange(locationExplored pokeapi.LocationAreaExplored, cfg *config) error {
	byteData, err := pokeapi.MarshalData(locationExplored.PokemonEncounters)
	if err != nil {
		return fmt.Errorf("error marshaling encounter data: %w", err)
	}
	cfg.Cache.Add("encounters-range", byteData)
	return nil
}
