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
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if command, exists := getCommands()[commandName]; exists {
			err := command.callback()
			if err != nil {
				fmt.Println("Error: ", err)
			}
			continue
		} else {
			fmt.Println("Unknown command. Type help for a list of available commands.")
			continue
		}
	}
}

func cleanInput(input string) []string {
	text := strings.ToLower(input)
	words := strings.Fields(text)

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
	}
}
