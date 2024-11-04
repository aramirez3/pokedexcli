package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("exit command")
	os.Exit(0)
	return nil
}

func commandHelp() error {
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

func commandMap() error {
	locationAreas, err := getLocationArea()
	if err != nil {
		return fmt.Errorf("error getting location-areas: %w", err)
	}
	fmt.Println("Get next 20 locations")
	return nil
}

func commandMapB() error {
	fmt.Println("Get previous 20 locations")
	return nil
}
