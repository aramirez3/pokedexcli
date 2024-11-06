package main

import (
	"time"

	"github.com/aramirez3/pokedexcli/internal/pokeapi"
	"github.com/aramirez3/pokedexcli/internal/pokecache"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(60 * time.Second)
	cfg := &config{
		pokeapiClient:   client,
		Cache:           cache,
		BaseCatchChance: 60,
	}
	startRepl(cfg)
}
