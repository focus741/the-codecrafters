# GO CONDITIONS

Conditional statements are used to perform different actions based on different conditions.

A condition can be either `true` or `false`.

Go supports the usual [comparison operators](https://www.w3schools.com/go/go_comparison_operators.php) from mathematics:

- Less than <
- Less than or equal <=
- Greater than >
- Greater than or equal >=
- Equal to ==
- Not equal to !=

Additionally, Go supports the usual [logical operators](https://www.w3schools.com/go/go_logical_operators.php):

- Logical AND &&
- Logical OR ||
- Logical NOT !

You can use these operators or their combinations to create conditions for different decisions.

Go has the following conditional statements:

- Use `if` to specify a block of code to be executed, if a specified condition is true
- Use `else` to specify a block of code to be executed, if the same condition is false
- Use `else if` to specify a new condition to test, if the first condition is false
- Use `switch` to specify many alternative blocks of code to be executed

## Go `if` Statement

Use the `if` statement to specify a block of Go code to be executed if a condition is `true`.

## Go `else` Statement

Use the `else` statement to specify a block of code to be executed if the condition is `false`.

## Go `else if` Statement

Use the `else if` statement to specify a new condition if the first condition is `false`.

## Go Nested if Statement

You can have `if` statements inside `if` statements, this is called a nested if.

## Go switch Statement

Use the `switch` statement to select one of many code blocks to be executed.

The `switch` statement in Go is similar to the ones in C, C++, Java, JavaScript, and PHP. The difference is that it only runs the matched case so it does not need a `break` statement.

## Go Multi-case switch Statement

It is possible to have multiple values for each `case` in the `switch` statement:
```
switch expression {
case x,y:
   // code block if expression is evaluated to x or y
case v,w:
   // code block if expression is evaluated to v or w
case z:
...
default:
   // code block if expression is not found in any cases
}
```

## Multi-case switch Example

The example below uses the weekday number to return different text:
```
package main
import ("fmt")

func main() {
   day := 5

   switch day {
   case 1,3,5:
    fmt.Println("Odd weekday")
   case 2,4:
     fmt.Println("Even weekday")
   case 6,7:
    fmt.Println("Weekend")
  default:
    fmt.Println("Invalid day of day number")
  }
}
```

---