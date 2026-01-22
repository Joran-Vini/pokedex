package main

import (
	"time"

	"github.com/Joran-Vini/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	conf := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.Pokemon),
	}
	startRepl(conf)
}
