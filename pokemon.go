package client

import (
	"context"
	"encoding/json"
	"net/http"
)

func (c *Client) GetPokemonByName(
	ctx context.Context,
	pokemonName string,
) (Pokemon, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.apiURL+"/api/v2/pokemon/"+pokemonName,
		nil,
	)

	if err != nil {
		return Pokemon{}, nil
	}

	req.Header.Add("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, PokemonFetchErr{
			Message:    err.Error(),
			StatusCode: -1,
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, PokemonFetchErr{
			Message:    "non-200 status code from the API",
			StatusCode: resp.StatusCode,
		}
	}

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return Pokemon{}, PokemonFetchErr{
			Message:    err.Error(),
			StatusCode: resp.StatusCode,
		}
	}

	return pokemon, nil
}
