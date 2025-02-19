package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Pizzu/pokedexcli/common"
	"github.com/Pizzu/pokedexcli/internal/pokeapi"
	"github.com/Pizzu/pokedexcli/internal/pokedex"
)

func main() {
	baseUrl := "https://pokeapi.co/api/v2/location-area?limit=20"
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeStorage := pokedex.NewMapStore()
	config := &config{pokeapiClient: pokeClient, pokedex: pokeStorage, Next: &baseUrl}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := common.CleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		var args []string

		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(config, args...)

			if err != nil {
				fmt.Println(err.Error())
			}
			continue
		} else {
			fmt.Println(baseUrl)
			fmt.Println("Unknown command")
			continue
		}
	}
}
