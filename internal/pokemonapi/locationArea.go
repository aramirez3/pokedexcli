package pokemonapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Location struct {
	id           int64
	name         string
	region       Region
	names        []Name
	game_indices []GenerationGameIndex
	area         []LocationArea
}

type LocationArea struct {
	id                     int64
	name                   string
	game_index             int64
	encounter_method_rates EncounterMethodRate
	location               Location
	names                  []Name
	pokemon_encounters     PokemonEncounter
}

type PokemonEncounter struct{}

type EncounterMethodRate struct{}

type Region struct {
	id   int64
	name string
}

type Name struct {
	name     string
	language Language
}

type Language struct {
	id       int64
	name     string
	official bool
	iso639   string
	iso3166  string
}

type GenerationGameIndex struct {
}

func getLocationArea(idOrName string) ([]Location, error) {
	endpoints := getEndPoints()
	url := endpoints["location"] + idOrName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []Location{}, fmt.Errorf("error building request: %w", err)
	}

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return []Location{}, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	locations := []Location{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return []Location{}, fmt.Errorf("error encoding json data: %w", err)
	}

	return locations, nil
}
