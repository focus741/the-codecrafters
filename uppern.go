package main

import (
	"fmt"
	"regexp"
	"strings"
)

func capN(input string) string {
	reCap := regexp.MustCompile(`\b(\w+) \(cap\)`)
	return reCap.ReplaceAllStringFunc(input, func(match string) string {
		parts := strings.Split(match, " ")
		word := parts[0]
		if len(word) == 0 {
			return match
		}
		return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	})
}

func main() {
	inputs := []string{
		"Welcome to the Brooklyn bridge (cap)",
		"the quick brown fox (cap) jumps",
		"no transformation here",
	}
	for _, s := range inputs {
		fmt.Println(capN(s))
	}
}
