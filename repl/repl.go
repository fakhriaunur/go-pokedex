package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fakhriaunur/go-pokedex/internal/commands"
)

func StartRepl(cfg *commands.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		if command, exists := commands.GetCommands()[commandName]; exists {
			err := command.Callback(cfg, args...)
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
