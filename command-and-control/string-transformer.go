package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Transformation struct {
	Command string
	Input   string
	Output  string
}

var History []Transformation

func showHistory() {
	start := len(History) - 5
	if start < 0 {
		start = 0
	}
	fmt.Println("--- Last 5 Transformations ---")
	for _, t := range History[start:] {
		fmt.Printf("Command: %s, In: %s, Out: %s\n", t.Command, t.Input, t.Output)
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	small := map[string]bool{"a": true, "an": true, "the": true, "of": true, "in": true, "on": true, "and": true, "to": true, "but": true, "or": true, "for": true, "nor": true, "at": true, "by": true, "up": true, "as": true, "is": true, "it": true}

	fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE  ")
	for {
		fmt.Println("Valid commands: *str upper <text>  *str lower <text>  *str cap <text>  *str title <text>  *str snake <text>  *str reverse <text>   *exit")

		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			main()
			return
		}

		parts := strings.Fields(input)
		varia := ""

		if parts[0] == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			return
		}
		for i := 0; i < 2; i++ {
			varia += strings.ToLower(parts[i])
			varia += " "
		}

		varia = strings.TrimSpace(varia)

		command := varia

		if command == "history" {
			showHistory()
			continue
		}

		if len(parts) < 2 {
			fmt.Println("ERROR:No command and text provided")
			continue
		}

		text := strings.Join(parts[2:], " ")

		var output string

		switch command {

		case "str upper":
			output = strings.ToUpper(text)

		case "str lower":
			output = strings.ToLower(text)

		case "str cap":
			words := strings.Fields(text)
			for i, w := range words {
				words[i] = strings.ToUpper(string(w[0])) + w[1:]
			}
			output = strings.Join(words, " ")

		case "str title":
			words := strings.Fields(text)
			for i, w := range words {
				lw := strings.ToLower(w)
				if i == 0 || !small[lw] {
					words[i] = strings.ToUpper(string(lw[0])) + lw[1:]
				} else {
					words[i] = lw
				}
			}
			output = strings.Join(words, " ")

		case "str snake":
			var r []rune
			for _, c := range text {
				if unicode.IsLetter(c) || unicode.IsDigit(c) {
					r = append(r, unicode.ToLower(c))
				} else if unicode.IsSpace(c) {
					r = append(r, '_')
				}
			}
			output = strings.Trim(strings.ReplaceAll(string(r), "__", "_"), "_")

		case "str reverse":
			words := strings.Fields(text)
			for i, w := range words {
				r := []rune(w)
				for x, y := 0, len(r)-1; x < y; x, y = x+1, y-1 {
					r[x], r[y] = r[y], r[x]
				}
				words[i] = string(r)
			}
			output = strings.Join(words, " ")

		default:
			fmt.Println(" Unknown command")
			continue

		}

		History = append(History, Transformation{Command: command, Input: text, Output: output})
		fmt.Println("→", output)
	}
}
