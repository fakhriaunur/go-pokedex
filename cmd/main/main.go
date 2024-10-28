package main

import (
	"time"

	"github.com/fakhriaunur/go-pokedex/internal/commands"
	"github.com/fakhriaunur/go-pokedex/internal/pokeapi"
	"github.com/fakhriaunur/go-pokedex/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &commands.Config{
		PokeapiClient: pokeClient,
	}
	repl.StartRepl(cfg)
}
