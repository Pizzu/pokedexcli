package common

import "strings"

func CleanInput(text string) []string {
	trimmedText := strings.TrimSpace(text)
	lowerCaseText := strings.ToLower(trimmedText)
	splittedSlice := strings.Split(lowerCaseText, " ")
	return splittedSlice
}
