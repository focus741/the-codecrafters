package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)


/*func showHistory() {
	hist := []string{}
	fmt.Println(hist)
}
*/
func main() {
	fmt.Println("Simple calculator: enter arguments")
	reader := bufio.NewReader(os.Stdin)
	//showHistory()

	for {
		fmt.Print(">> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input", err)
			continue
		}

		line = strings.TrimSpace(line)
		parts := strings.Fields(line)

		if len(parts) == 0 {
			fmt.Println("Type something...")
			continue
		}

		command := strings.ToLower(parts[0])
		if command == "quit" {
			fmt.Println("Goodbye")
			return
		}

		if command == "help" {
			dispayHelp()
			continue
		}

		/*if command == "history" {
			showHistory()
			continue
		}*/

		if len(parts) != 3 {
			fmt.Println("Error: three arguments required <operation> <a> <b>")
			continue
		}

		num1, err1 := strconv.ParseFloat(parts[1], 64)
		num2, err2 := strconv.ParseFloat(parts[2], 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Error, one of both of your numbers is invalid\nExample add 5 2")
			continue
		}

		
		//hist := []string{}

		switch command {
		case "add":
			/*fmt.Println("Result:", num1+num2)
			a := num1 + num2
			aInt := int64(a)
			aStr := strconv.FormatInt(aInt, 10)
			hist := append(hist, command+parts[1]+parts[2]+"Result: "+aStr)*/
		case "sub":
			fmt.Println("Result:", num1-num2)
		case "mul":
			fmt.Println("Result:", num1*num2)
		case "pow":
			fmt.Println("Result:", math.Pow(num1, num2))
		case "mod":
			fmt.Println("Result:", math.Mod(num1, num2))
		case "div":
			if num2 == 0 {
				fmt.Println("Division by zero is not allowed")
				continue
			}
			fmt.Println("Result:", num1/num2)
		default:
			fmt.Println("unknown command\ntry <help>")
		}
	}
}

func dispayHelp() {
	fmt.Println("add <a> <b>   → addition")
	fmt.Println("sub <a> <b>   → subtraction")
	fmt.Println("mul <a> <b>   → multiplication")
	fmt.Println("div <a> <b>   → division")
}