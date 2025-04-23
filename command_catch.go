package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area specified")
	}

	pokemonName := args[0]
	resp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(resp.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("Failed to Catch! %s", pokemonName)
	}

	fmt.Println("Caught!")
	fmt.Println("Congratulations!")
	fmt.Println("You have a new Pokemon!")
	fmt.Println("Enjoy your new companion!")
	return nil
}
