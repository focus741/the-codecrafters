package main

import (
	"fmt"
	"strings"
	"unicode"
)

func fixQuotes(s string) string {
	var result strings.Builder
	n := len(s)
	inQuote := false
	for i := 0; i < n; i++ {
		if s[i] == '\'' {
			if !inQuote {
				inQuote = true
				result.WriteByte('\'')
				j := i + 1
				for j < n && s[j] == ' ' {
					j++
				}
				i = j - 1
			} else {
				inQuote = false
				resStr := result.String()
				if len(resStr) > 0 && resStr[len(resStr)-1] == ' ' {
					result.Reset()
					result.WriteString(resStr[:len(resStr)-1])
				}
				result.WriteByte('\'')
			}
		} else {
			if unicode.IsSpace(rune(s[i])) {
				if i+1 < n && !unicode.IsSpace(rune(s[i+1])) {
					result.WriteByte(' ')
				}
			} else {
				result.WriteByte(s[i])
			}
		}
	}
	return strings.TrimSpace(result.String())
}

func main() {
	input1 := "  I am exactly how they describe me: ' awesome    '       "
	input2 := "   As Agbaji alexander said: ' I       am the most well-known smartest guy in the world    .     '"

	fmt.Println(fixQuotes(input1))
	fmt.Println(fixQuotes(input2))
}
