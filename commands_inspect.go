package main

import (
	"fmt"

	"github.com/aramirez3/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config) error {
	if len(cfg.Words) == 0 {
		return fmt.Errorf("no Pokemon name was entered")
	}
	name := cfg.Words[0]
	fmt.Printf("Inspect %s\n", name)
	byteData, found := cfg.Cache.Get(name)
	pokemon := pokeapi.Pokemon{}
	pokeapi.UnmarshalData(byteData, &pokemon)
	if found {
		fmt.Printf("  Name: %s\n", pokemon.Name)
	} else {
		fmt.Println("You have not caught that Pokemon")
	}
	return nil
}
