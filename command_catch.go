package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, args ...string) error {
	pokemonName := args[0]

	if len(args) == 0 {
		return fmt.Errorf("Please provide a Pokemon name to catch.")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemon, err := conf.pokeapiClient.PokemonGet(pokemonName)
	if err != nil {
		return err
	}
	chance := rand.Intn(pokemon.BaseExperience)
	if chance < 40 {
		fmt.Printf("Caught %s!\n", pokemon.Name)
		conf.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
