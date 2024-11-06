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

func MarshalData[T any](fromType T) ([]byte, error) {
	jsonData, err := json.Marshal(fromType)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data: %w", err)
	}
	return jsonData, nil
}
