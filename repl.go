package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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
			cmd.callback()
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	Next     *string
	Previous *string
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
		"map": {
			name:        "map",
			description: "Display the names of the next 20 locations in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the previous 20 locations in the Pokemon world",
			callback:    commandMapB,
		},
	}
}
