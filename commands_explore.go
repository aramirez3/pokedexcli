package main

import (
	"fmt"
)

func commandExplore(cfg *config) error {
	areaName := ""
	if len(cfg.Words) > 0 {
		areaName = cfg.Words[0]
	} else {
		fmt.Println("explore requires an area-name")
		return nil
	}

	locationExplored, _ := cfg.pokeapiClient.ExploreLocation(&areaName)
	fmt.Println(locationExplored)
	fmt.Printf("Exploring %s...\n", locationExplored.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range locationExplored.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
