package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(capN("i remember my name, like yesterday (cap, 6)"))
}

func capN(s string) string {
	open_Tag := strings.Index(s, "(")
	close_Tag := strings.Index(s, ")")
	Tag := s[open_Tag : close_Tag+1]

	beforeTag := s[:open_Tag]
	words := strings.Fields(beforeTag)

	n := 0
	for i := 0; i < len(Tag)-1; i++ {
		if strings.ContainsAny(string(Tag[i]), "123456789") {
			n, _ = strconv.Atoi(string(Tag[i]))
		}
	}

	for i := len(words) - n; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	words = append(words, s[close_Tag+1:])

	return strings.Join(words, " ")
}
