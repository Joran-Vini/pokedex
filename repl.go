package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(conf)
		if err != nil {
			fmt.Println(err)
		}
	}
}
func cleanInput(text string) []string {

	to_lower := strings.ToLower(text)
	trimmed := strings.TrimSpace(to_lower)
	words := strings.Fields(trimmed)

	return words
}
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the map of the region",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of the map",
			callback:    commandMapBack,
		},
	}
}
