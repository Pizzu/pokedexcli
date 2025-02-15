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
		for scanner.Scan() {
			firstWord := cleanInput(scanner.Text())
			fmt.Printf("Your command was: %s \n", firstWord[0])
			break
		}
	}
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	lowerCaseText := strings.ToLower(trimmedText)
	splittedSlice := strings.Split(lowerCaseText, " ")
	return splittedSlice
}
