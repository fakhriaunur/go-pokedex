package commands

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.PokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!...\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!...\n", pokemon.Name)
	cfg.CaughtPokemon[pokemon.Name] = pokemon

	return nil
}
