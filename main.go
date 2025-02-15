package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Pizzu/pokedexcli/common"
)

func main() {
	baseUrl := "https://pokeapi.co/api/v2/location-area?limit=20"
	config := &config{Next: &baseUrl, Previous: nil}
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
			fmt.Println("Unknown command")
			continue
		}
	}
}
