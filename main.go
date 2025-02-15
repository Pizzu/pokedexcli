package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		fmt.Printf("Your command was: %s \n", commandName)
	}
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	lowerCaseText := strings.ToLower(trimmedText)
	splittedSlice := strings.Split(lowerCaseText, " ")
	return splittedSlice
}
