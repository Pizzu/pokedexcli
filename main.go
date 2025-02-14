package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	splittedSlice := strings.Split(trimmedText, " ")
	return splittedSlice
}
