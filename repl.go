package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aramirez3/pokedexcli/internal/pokeapi"
	"github.com/aramirez3/pokedexcli/internal/pokecache"
)

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()
		text := reader.Text()
		words := strings.Fields(text)

		commands := getCommands()
		cmd, ok := commands[words[0]]
		if !ok {
			fmt.Printf("unknown command: %s\n", words[0])
			continue
		} else {
			cfg.Words = words[1:]
			cmd.callback(cfg)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	Cache         *pokecache.Cache
	Words         []string
	Next          *string
	Previous      *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"catch": {
			name:        "catch",
			description: "Throw a Pokeball at a Pokemon and attempt to catch them!",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore",
			description: "Displays Pokenmon characters found within a given area",
			callback:    commandExplore,
		},
		"mapf": {
			name:        "mapf",
			description: "Display the names of the next 20 locations in the Pokemon world",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the previous 20 locations in the Pokemon world",
			callback:    commandMapB,
		},
	}
}
