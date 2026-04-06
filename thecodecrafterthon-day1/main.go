package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero\n")
	}
	return a / b, nil
}

func main() {
	for {
		var a, b int
		fmt.Println("Input first number:")
		fmt.Scanln(&a)

		fmt.Println("Input second number:")
		fmt.Scanln(&b)

		var op string

		fmt.Println("Expression = add / sub / mul / div / help / quit:")
		fmt.Scanln(&op)

		switch op {
		case "add":
			fmt.Println("Result:", add(a, b))

		case "sub":
			fmt.Println("Result:", sub(a, b))

		case "mul":
			fmt.Println("Result:", mul(a, b))

		case "div":
			result, err := div(a, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Result:", result)
			}

		case "help":
			fmt.Println("Supported commands: add / sub / mul / div / help / quit")

		case "quit":
			fmt.Println("Goodbye King Focus")
			return

		default:
			fmt.Println("Invalid Expression!!!")
		}
	}
}
