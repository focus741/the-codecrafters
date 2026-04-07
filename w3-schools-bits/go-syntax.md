# GO SYNTAX

## A Go file consists of the following parts:

- Package declaration
- Import packages
- Functions
- Statements and expressions

```
package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")
}
```

## Example explanation

**Line 1:** In **GO** every program is a part of a package. We define this using ```package``` keyword. So this example belong to the ```main``` package.

**Line 2:** ```import ("fmt")``` lets us import files included in the ```fmt``` paackage.

**Line 3:** A blank line. **GO** ignores white space and white spaces makes our codes more readable.

**Line 4:** ```func main() {}``` is a function. Any code inside its curly brackets `{}` will be executed.

**Line 5:** `fmt.Println()` is a function made available from the `fmt` package. It is used to print text. In our example it will print "Hello World!".

## Go Statements

`fmt.Priintln("Hello World")` is a ststement.

In **GO** statements are separated by the ending a line (hitting the ENTER KEY) or by a simicolon ";".

## Go Compact Code

You can write more compact code using the simicolon. But its make codes hard to read, example shown below:

```package main; import ("fmt"); func main() { fmt.Println("Hello World!");}```
