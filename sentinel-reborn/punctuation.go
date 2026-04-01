package main

import (
	"fmt"
	"strings"
)

func fixQuotes(s string) string {
	fields := strings.Fields(s)
	s = strings.Join(fields, " ")

	var result strings.Builder
	inQuote := false

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == '\'' {
			if !inQuote {
				inQuote = true
				result.WriteByte('\'')
				if i+1 < len(s) && s[i+1] == ' ' {
					i++
				}
			} else {
				inQuote = false
				resStr := result.String()
				if len(resStr) > 0 && resStr[len(resStr)-1] == ' ' {
					result.Reset()
					result.WriteString(strings.TrimSuffix(resStr, " "))
				}
				result.WriteByte('\'')
			}
		} else if char == '.' {
			resStr := result.String()
			if len(resStr) > 0 && resStr[len(resStr)-1] == ' ' {
				result.Reset()
				result.WriteString(strings.TrimSuffix(resStr, " "))
			}
			result.WriteByte('.')
		} else {
			result.WriteByte(char)
		}
	}

	return result.String()
}

func main() {
	input1 := "  I am exactly how they describe me: ' awesome    '       "
	input2 := "   As Agbaji alexander said: ' I       am the most well-known smartest guy in the world    .     '"

	fmt.Println(fixQuotes(input1))
	fmt.Println(fixQuotes(input2))
}
