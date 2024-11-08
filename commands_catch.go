package main

import (
	"fmt"
	"math/rand"

	"github.com/aramirez3/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config) error {
	if len(cfg.Words) == 0 {
		return fmt.Errorf("no Pokemon name was entered")
	}
	name := &cfg.Words[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return fmt.Errorf("error getting pokemon: %w", err)
	}
	success := pokemonIsCaught(cfg, pokemon.BaseExperience)
	fmt.Printf("Throwing at ball at %s...\n", pokemon.Name)
	if success {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		addPokemonToCache(cfg, pokemon)
	} else {
		fmt.Printf("%s got got away!\n", pokemon.Name)
	}
	return nil
}

func pokemonIsCaught(cfg *config, baseExp int64) bool {
	threshold := float64(200)
	adjustedChance := float64(cfg.BaseCatchChance) - (float64(baseExp) / threshold)
	randNumber := rand.Intn(100)
	return adjustedChance < float64(randNumber)
}

func addPokemonToCache(cfg *config, pokemon pokeapi.Pokemon) error {
	cfg.Pokedex = append(cfg.Pokedex, pokemon.Name)
	byteData, err := pokeapi.MarshalData(pokemon)
	if err != nil {
		return fmt.Errorf("error marshaling pokemon data: %w", err)
	}
	cfg.Cache.Add(pokemon.Name, byteData)
	return nil
}
