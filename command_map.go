package main

import (
	"fmt"
)

func commandMap(conf *config) error {
	locationsResp, err := conf.pokeapiClient.ListLocations(conf.next)
	if err != nil {
		return err
	}
	conf.next = locationsResp.Next
	conf.previous = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}
func commandMapBack(conf *config) error {
	if conf.previous == nil {
		fmt.Println("No previous page available.")
		return nil
	}
	locationsResp, err := conf.pokeapiClient.ListLocations(conf.previous)
	if err != nil {
		return err
	}
	conf.next = locationsResp.Next
	conf.previous = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}
