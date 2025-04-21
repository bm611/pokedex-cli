package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Unknown command:", commandName)
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Show this help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the REPL",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Show location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous location areas",
			callback:    callbackMapb,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
