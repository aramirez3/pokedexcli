package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()
		text := reader.Text()
		println(text)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}