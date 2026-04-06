package main

import (
	"fmt"
	"strings"
)

func transformString(input string) string {
	// Reverse the words in the string
	words := strings.Fields(input)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func allCap(input string) string {
	return strings.ToUpper(input)
}

func allLow(input string) string {
	return strings.ToLower(input)
}

func trimSpace(input string) string {
	return strings.TrimSpace(input)
}

func validateInput(input string) bool {
	return input != ""
}

func titleCase(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

func main() {
	input := "hello, world!"
	if !validateInput(input) {
		fmt.Println("Input cannot be empty.")
		return
	}
	input = strings.TrimSpace(input)
	transformed := transformString(input)
	fmt.Println(transformed)

	capitalized := allCap(input)
	fmt.Println(capitalized)

	lowercased := allLow(input)
	fmt.Println(lowercased)

	trimmed := trimSpace(input)
	fmt.Println(trimmed)
	titleCased := titleCase(input)
	fmt.Println(titleCased)
}
