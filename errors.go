package client

import (
	"errors"
	"fmt"
)

var (
	ErrFetchingPokemon = errors.New("failed to fetch pokemon")
)

type PokemonFetchErr struct {
	StatusCode int
	Message    string
}

func (p PokemonFetchErr) Error() string {
	return fmt.Sprintf("failed to fetch pokemon: %s with status code: %d", p.Message, p.StatusCode)
}
