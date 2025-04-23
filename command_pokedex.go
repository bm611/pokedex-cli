package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	for _, v := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", v.Name)
	}
	return nil
}
