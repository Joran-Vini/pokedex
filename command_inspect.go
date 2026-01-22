package main

import (
	"fmt"
)

func commandInspect(conf *config, args ...string) error {
	pokemonName := args[0]

	if len(args) == 0 {
		return fmt.Errorf("Please provide a Pokemon name to inspect.")
	}

	pokemon, exists := conf.pokedex[pokemonName]
	if !exists {
		return fmt.Errorf("You don't have %s in your Pokedex.", pokemonName)
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}
