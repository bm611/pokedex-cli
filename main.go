package main

import "github.com/bm611/pokedex-cli/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	// pointer because it can be nil
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient:       pokeapi.NewClient(),
		nextLocationAreaURL: nil,
		prevLocationAreaURL: nil,
	}
	startRepl(&cfg)

}
