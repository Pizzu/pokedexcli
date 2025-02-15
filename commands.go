package main

import (
	"fmt"
	"os"

	"github.com/Pizzu/pokedexcli/api"
)

type config struct {
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
		"map": {
			name:        "map",
			description: "Display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "Display the previous 20 locations",
			callback:    commandMapBack,
		},
	}
}

func commandHelp(config *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(config *config) error {
	if config.Next != nil {
		locationDTO, err := api.GetAllLocationArea(*config.Next)

		if err != nil {
			return err
		}

		for _, result := range locationDTO.Results {
			fmt.Println(result.Name)
		}

		config.Next = locationDTO.Next
		config.Previous = locationDTO.Previous
	} else {
		fmt.Println("you're on the last page")
	}

	return nil
}

func commandMapBack(config *config) error {
	if config.Previous != nil {
		locationDTO, err := api.GetAllLocationArea(*config.Previous)

		if err != nil {
			return err
		}

		for _, result := range locationDTO.Results {
			fmt.Println(result.Name)
		}

		config.Next = locationDTO.Next
		config.Previous = locationDTO.Previous

	} else {
		fmt.Println("you're on the first page")
	}

	return nil
}
