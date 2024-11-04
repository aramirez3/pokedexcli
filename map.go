package main

import (
	"fmt"
)

func commandMapF(cfg *config) error {
	locationAreas, err := cfg.pokeapiClient.GetLocations(*cfg.next)
	if err != nil {
		return fmt.Errorf("error getting location-areas: %w", err)
	}
	fmt.Println("Get next 20 locations")
	fmt.Println(locationAreas)
	return nil
}

func commandMapB(cfg *config) error {
	fmt.Println("Get previous 20 locations")
	return nil
}
