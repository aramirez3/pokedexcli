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
	locationExplored, _ := cfg.pokeapiClient.ExploreLocation(&areaName)
	byteData, _ := pokeapi.MarshalData(locationExplored)
	cfg.Cache.Add(areaName, byteData)
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
