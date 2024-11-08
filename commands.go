package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("exit command")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
