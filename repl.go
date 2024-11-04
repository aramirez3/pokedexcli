package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aramirez3/pokedexcli/internal/pokeapi"
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
	next          *string
	previous      *string
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
