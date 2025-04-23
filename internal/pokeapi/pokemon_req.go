package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	fullURL := baseURL + "pokemon/" + pokemonName
	var pokemon Pokemon

	// check the cache
	if data, ok := c.cache.Get(fullURL); ok {
		fmt.Println("Cache hit")
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	// Make the API request
	resp, err := c.httpClient.Get(fullURL)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)
	return pokemon, nil
}
