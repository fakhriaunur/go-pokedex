package commands

import "fmt"

func commandPokedex(cfg *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	pokedex := cfg.CaughtPokemon
	if len(pokedex) == 0 {
		fmt.Println("You haven't caught any pokemon yet!")
		return nil
	}

	for _, poke := range pokedex {
		fmt.Printf("  - %s\n", poke.Name)
	}

	return nil
}
