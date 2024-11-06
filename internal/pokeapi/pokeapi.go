package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	BaseUrl          = "https://pokeapi.co/api/v2"
	LocationEndpoint = "location-area"
	PokemonEndpoint  = "pokemon"
)

func (c *Client) GetUrl(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	response := []byte{}
	if err != nil {
		return response, fmt.Errorf("error making request: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return response, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return response, fmt.Errorf("error reading data: %w", err)
	}
	return data, nil
}

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
