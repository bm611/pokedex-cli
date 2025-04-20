package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]
		switch command {
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  help - Show this help message")
			fmt.Println("  exit - Exit the REPL")
		case "exit":
			fmt.Println("Exiting REPL...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command:", command)
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
