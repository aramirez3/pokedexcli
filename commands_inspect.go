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
		printPokemonStats(pokemon)
	} else {
		fmt.Println("You have not caught that Pokemon")
	}
	return nil
}

func printPokemonStats(p pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range p.Stats {
		fmt.Printf("    - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, stat := range p.Types {
		fmt.Printf("    - %s\n", stat.Type.Name)
	}
}
