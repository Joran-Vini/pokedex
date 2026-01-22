package main

import "github.com/Joran-Vini/pokedex/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
}
