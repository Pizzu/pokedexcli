package main

import (
	"strings"
)

func main() {
	startRepl()
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	lowerCaseText := strings.ToLower(trimmedText)
	splittedSlice := strings.Split(lowerCaseText, " ")
	return splittedSlice
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
