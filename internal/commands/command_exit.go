package commands

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config) error {
	fmt.Println("Exiting the Pokedex...")
	os.Exit(0)

	return nil
}
