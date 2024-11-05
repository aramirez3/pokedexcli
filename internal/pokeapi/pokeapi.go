package pokeapi

import (
	"encoding/json"
	"fmt"
)

const (
	baseUrl      = "https://pokeapi.co/api/v2"
	locationArea = "location-area"
)

func UnmarshalData[T any](data []byte, toType T) error {
	err := json.Unmarshal(data, &toType)
	if err != nil {
		return fmt.Errorf("error unmarshaling data: %w", err)
	}
	return nil
}
