package main

import (
	"fmt"
)

func commandPokedex(cfg *config) error {
	pokedex := cfg.Pokedex
	if len(pokedex) > 0 {
		fmt.Println("Your Pokedex:")
		for _, name := range cfg.Pokedex {
			fmt.Printf("    - %s\n", name)
		}
	} else {
		fmt.Println("there are no items in your Pokedex")
	}
	return nil
}
