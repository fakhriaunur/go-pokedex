package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Exiting the Pokedex...")
	os.Exit(0)

	return nil
}