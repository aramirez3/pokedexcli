package main

import (
	"fmt"
)

func commandMapF(cfg *config) error {
	locationAreas, err := cfg.pokeapiClient.GetLocations(cfg.Next)
	if err != nil {
		return fmt.Errorf("error getting location-areas: %w\n", err)
	}
	cfg.Next = locationAreas.Next
	cfg.Previous = locationAreas.Previous
	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	fmt.Println("Get previous 20 locations")
	return nil
}
