package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Pizzu/pokedexcli/common"
	"github.com/Pizzu/pokedexcli/internal/pokeapi"
	"github.com/Pizzu/pokedexcli/internal/pokedex"
)

type config struct {
	pokeapiClient pokeapi.Client
	pokedex       pokedex.PokemonStore
	Next          *string
	Previous      *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"explore": {
			name:        "explore",
			description: "list of all pokemon located in a specific location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "try to catch a pokemon e.g. catch charizard",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "get more info about a specific captured pokemon",
			callback:    commandInspect,
		},
	}
}

func commandHelp(config *config, arg ...string) error {
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

func commandExit(config *config, arg ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(config *config, arg ...string) error {
	if config.Next != nil {
		locationDTO, err := config.pokeapiClient.ListLocations(config.Next)

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

func commandMapBack(config *config, arg ...string) error {
	if config.Previous != nil {
		locationDTO, err := config.pokeapiClient.ListLocations(config.Previous)

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

func commandExplore(config *config, args ...string) error {
	if args == nil {
		return errors.New("please provide a location area")
	}

	pokemonLocationDTO, err := config.pokeapiClient.ListPokemonsWithinLocation(args[0])

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	for _, data := range pokemonLocationDTO.PokemonEncounters {
		pokemonName := data.Pokemon.Name
		fmt.Printf("- %s\n", pokemonName)
	}

	return nil
}

func commandCatch(config *config, args ...string) error {
	if args == nil {
		return errors.New("please provide a pokemon name")
	}

	pokemonDTO, err := config.pokeapiClient.GetPokemon(args[0])

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonDTO.Name)

	isCatched := common.AttempCatch(pokemonDTO.BaseExperience)

	if isCatched {
		err = config.pokedex.Add(pokemonDTO)

		if err != nil {
			return err
		}

		fmt.Printf("%s was caught!\n", pokemonDTO.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemonDTO.Name)
	}

	return nil
}

func commandInspect(config *config, args ...string) error {
	if args == nil {
		return errors.New("please provide a pokemon name")
	}

	pokemonDTO, err := config.pokedex.GetPokemon(args[0])

	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", pokemonDTO.Name)
	fmt.Printf("Height: %v\n", pokemonDTO.Height)
	fmt.Printf("Weight: %v\n", pokemonDTO.Weight)

	fmt.Println("Types:")

	for _, value := range pokemonDTO.Types {
		fmt.Printf(" - %s\n", value.Type.Name)
	}

	return nil
}
