package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*BASE CONVERTER */

func hexToDec(hext string) int64 {
	val, err := strconv.ParseInt(hext, 16, 64)
	if err != nil {
		fmt.Println("Invalid hex input")
	}
	return val
}

func binToDec(bina string) int64 {
	val, err := strconv.ParseInt(bina, 2, 64)
	if err != nil {
		fmt.Println("Invalid binary input")
	}
	return val
}

func decConvert(deci string) (string, string) {
	d, err := strconv.ParseInt(deci, 10, 64)
	if err != nil {
		fmt.Println("Invalid decimal input")
	}
	return strconv.FormatInt(d, 2), strings.ToUpper(strconv.FormatInt(d, 16))
}

func baseConverter() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- BASE CONVERTER ---")
		fmt.Println("Commands: hex <value>, bin <value>, dec <value>, back")

		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(strings.TrimSpace(line))

		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]

		if cmd == "back" {
			return
		}

		if len(parts) < 2 {
			fmt.Println("Provide a value")
			continue
		}

		input := parts[1]

		switch cmd {
		case "hex":
			fmt.Println("Decimal:", hexToDec(input))
		case "bin":
			fmt.Println("Decimal:", binToDec(input))
		case "dec":
			b, h := decConvert(input)
			fmt.Println("Binary:", b)
			fmt.Println("Hex:", h)
		default:
			fmt.Println("Unknown command")
		}
	}
}

/*CALCULATOR*/

func calculator() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- CALCULATOR ---")
		fmt.Println("Format: add 5 2 | sub | mul | div | pow | mod | back")

		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(strings.TrimSpace(line))

		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]

		if cmd == "back" {
			return
		}

		if len(parts) != 3 {
			fmt.Println("Usage: <operation> <a> <b>")
			continue
		}

		a, err1 := strconv.ParseFloat(parts[1], 64)
		b, err2 := strconv.ParseFloat(parts[2], 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid numbers")
			continue
		}

		switch cmd {
		case "add":
			fmt.Println("Result:", a+b)
		case "sub":
			fmt.Println("Result:", a-b)
		case "mul":
			fmt.Println("Result:", a*b)
		case "div":
			if b == 0 {
				fmt.Println("Cannot divide by zero")
				continue
			}
			fmt.Println("Result:", a/b)
		case "pow":
			fmt.Println("Result:", math.Pow(a, b))
		case "mod":
			fmt.Println("Result:", math.Mod(a, b))
		default:
			fmt.Println("Unknown command")
		}
	}
}

/* STRING TRANSFORMER */

type Transformation struct {
	Command string
	Input   string
	Output  string
}

var history []Transformation

func showHistory() {
	fmt.Println("--- Last Transformations ---")
	for _, h := range history {
		fmt.Printf("%s | %s → %s\n", h.Command, h.Input, h.Output)
	}
}

func stringTransformer() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- STRING TRANSFORMER ---")
		fmt.Println("Commands: upper, lower, cap, reverse, snake, history, back")
		fmt.Print("> ")

		if !scanner.Scan() {
			return
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		cmd := parts[0]

		if cmd == "back" {
			return
		}

		if cmd == "history" {
			showHistory()
			continue
		}

		if len(parts) < 2 {
			fmt.Println("Provide text")
			continue
		}

		text := strings.Join(parts[1:], " ")
		var output string

		switch cmd {
		case "upper":
			output = strings.ToUpper(text)

		case "lower":
			output = strings.ToLower(text)

		case "cap":
			words := strings.Fields(text)
			for i, w := range words {
				words[i] = strings.ToUpper(string(w[0])) + w[1:]
			}
			output = strings.Join(words, " ")

		case "reverse":
			r := []rune(text)
			for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
				r[i], r[j] = r[j], r[i]
			}
			output = string(r)

		case "snake":
			var r []rune
			for _, c := range text {
				if unicode.IsLetter(c) || unicode.IsDigit(c) {
					r = append(r, unicode.ToLower(c))
				} else if unicode.IsSpace(c) {
					r = append(r, '_')
				}
			}
			output = string(r)

		default:
			fmt.Println("Unknown command")
			continue
		}

		history = append(history, Transformation{cmd, text, output})
		fmt.Println("→", output)
	}
}

/* MAIN MENU */

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MULTI TOOL CLI ===")
		fmt.Println("1. base  → Base Converter")
		fmt.Println("2. calc  → Calculator")
		fmt.Println("3. str   → String Transformer")
		fmt.Println("4. quit")

		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		cmd := strings.TrimSpace(input)

		switch cmd {
		case "base":
			baseConverter()
		case "calc":
			calculator()
		case "str":
			stringTransformer()
		case "quit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
