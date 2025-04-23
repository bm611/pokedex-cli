package main

import (
	"time"

	"github.com/bm611/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	// pointer because it can be nil
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient:       pokeapi.NewClient(time.Hour),
		nextLocationAreaURL: nil,
		prevLocationAreaURL: nil,
		caughtPokemon:       make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
