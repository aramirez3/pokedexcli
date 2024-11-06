package main

import (
	"fmt"
)

func commandCatch(cfg *config) error {
	fmt.Println("Catch command")
	cfg.pokeapiClient.CatchPokemon(&cfg.Words[0])
	return nil
}
