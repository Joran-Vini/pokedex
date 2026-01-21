package main

import (
	"strings"
)

func cleanInput(text string) []string {

	to_lower := strings.ToLower(text)
	trimmed := strings.TrimSpace(to_lower)
	words := strings.Fields(trimmed)
	

	return words
}
