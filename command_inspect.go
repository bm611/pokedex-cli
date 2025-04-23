package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemonspecified")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("pokemon not caught yet!")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Type: %s\n", pokemon.Types[0].Type.Name)
	fmt.Printf("HP: %d/%d\n", pokemon.Stats[0].BaseStat, pokemon.Stats[0].BaseStat)
	fmt.Printf("Attack: %d\n", pokemon.Stats[1].BaseStat)
	fmt.Printf("Defense: %d\n", pokemon.Stats[2].BaseStat)
	fmt.Printf("Speed: %d\n", pokemon.Stats[5].BaseStat)
	fmt.Printf("Special Attack: %d\n", pokemon.Stats[3].BaseStat)
	fmt.Printf("Special Defense: %d\n", pokemon.Stats[4].BaseStat)
	fmt.Printf("Experience: %d\n", pokemon.BaseExperience)

	return nil
}
