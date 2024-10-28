package commands

import "github.com/fakhriaunur/go-pokedex/internal/pokeapi"

type Config struct {
	PokeapiClient    pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
}
type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exits the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Displays location-area",
			Callback:    commandMapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Navigates to previous location-area",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore <location>",
			Description: "Explore a specific location",
			Callback:    commandExplore,
		},
	}
}
