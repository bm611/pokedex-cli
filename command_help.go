package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("Welcome to Pokedex CLI!")
	fmt.Println("Available commands:")
	availableCommands := getCommands()
	for _, command := range availableCommands {
		fmt.Printf("  %s - %s\n", command.name, command.description)
	}
	return nil
}
