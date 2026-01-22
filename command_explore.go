package main

import (
	"fmt"
)

func commandExplore(conf *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide a location name to explore.")
	}
	locationName := args[0]
	location, err := conf.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s area...\n", locationName)
	fmt.Println("Found Pokémon: ")
	for _, poke := range location.PokemonEncounters {
		fmt.Printf("- %s\n", poke.Pokemon.Name)
	}
	return nil
}
