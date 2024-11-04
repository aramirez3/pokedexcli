package main

import (
	"time"

	"github.com/aramirez3/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: client,
	}
	startRepl(cfg)
}
