package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Pizzu/pokedexcli/common"
	"github.com/Pizzu/pokedexcli/internal/pokeapi"
)

func main() {
	baseUrl := "https://pokeapi.co/api/v2/location-area?limit=20"
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &config{pokeapiClient: pokeClient, Next: &baseUrl}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := common.CleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(config)

			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println(baseUrl)
			fmt.Println("Unknown command")
			continue
		}
	}
}
