# GO FUNCTIONS

A function is a block of statements that can be used repeatedly in a program, function will not execute automatically when a page loads, function will be executed by a call to the function.

## Create a Function

To create (often referred to as declare) a function, do the following:

Use the `func` keyword.
Specify a name for the function, followed by parentheses ().
Finally, add code that defines what the function should do, inside curly braces {}.

## Call a Function

Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.

## Naming Rules for Go Functions

- A function name must start with a letter
- A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
- Function names are case-sensitive
- A function name cannot contain spaces
- If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used

## Go Function Parameters and Arguments

### Parameters and Arguments

Information can be passed to functions as a parameter. Parameters act as variables inside the function.

Parameters and their types are specified after the function name, inside the parentheses. You can add as many parameters as you want, just separate them with a comma.

### Function With Parameter Example

The following example has a function with one parameter (`fname`) of type `string`. When the familyName() function is called, we also pass along a name (e.g. Liam), and the name is used inside the function, which outputs several different first names, but an equal last name.

### Multiple Parameters

Inside the function, you can add as many parameters as you want.

## Go Function Returns

### Return Values

If you want the function to return a value, you need to define the data type of the return value (such as `int`, `string`, etc), and also use the `return` keyword inside the function.

## Go Recursion Functions

Go accepts recursion functions. A function is recursive if it calls itself and reaches a stop condition.

In the following example, `testcount()` is a function that calls itself. We use the `x` variable as the data, which increments with 1 (`x + 1`) every time we recurse. The recursion ends when the `x` variable equals to 11 (`x == 11`). 
```
package main
import ("fmt")

func testcount(x int) int {
  if x == 11 {
    return 0
  }
  fmt.Println(x)
  return testcount(x + 1)
}

func main(){
  testcount(1)
}
```

