# GO VARIABLES

---

## Go Variable Types

In Go, there are different types of variables, which are; `int`, `float32`, `string` and `bool`.
`int` - for storing intergers, such as 123 or -123.
`float32` - for storing floating point numbers, with decimal, such as 19.9 or -19.9.
`string` - for storing text, such as "Hello World". String values are surrounded by double quotes.
`bool` - for storing values with two states: true or false.

---

## Declaring (Creating) Variables

---

In Go, there are two types of ways to declare a variable:

- With the `var` keyword:
`var variablename type = value`

- With the `:=` sign:
`variablename := value`

---

## Variable Declaration With Initial Value

If the value of a variable is known from the start, you can declare the variable and assign a value to it onone line:

```
package main
import ("fmt")
func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred
  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
}
```

---

## Variable Declaration Without Initial Value

In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type:

```
package main
import ("fmt")
func main() {
  var a string
  var b int
  var c bool
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
```

---

## Value Assignment After Declaration

It is possible to assign a value to a variable after it is declared. This is helpful for cases the value is not initially known.

---

## Difference Between var and :=

`var` can be used inside and outside of a functions while `:=` can only be used inside functions.

`var` variable declaration and value assignment can be done separately while in `:=` variable declaration and value assignment cannot be done separately(must be done in the same line)

---

## Go Multiple Variable Declaration

In Go, it is possible to declare multiple variables on the same line.
e.g `var a, b, c, d int = 1, 3, 5, 7`

If the type keyword is not specified, you can declare different types of variables on the same line:
```
var a, b = 6, "Hello"
  c, d := 7, "World!"
  ```

  ---

  ## Go Variable Declaration in a Block

  Multiple variable declarations can also be grouped together into a block for greater readability:
  ```
  var (
     a int
     b int = 1
     c string = "hello"
   )
   ```

  ---

  ## Go Variable Naming Rules

  A variable can have a short name (like x and y) or a more descritive name like age, price, carname e.t.c...

  Go variable naming rules:

- A variable name must start with a letter or an underscore character (_)
- A variable name cannot start with a digit
- A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
- Variable names are case-sensitive (age, Age and AGE are three different variables)
- There is no limit on the length of the variable name
- A variable name cannot contain spaces
- The variable name cannot be any Go keywords

---

## Multi-Word Variable Names

Variables names with more than one word can be difficult to read. There are several techniques you can use to make them more readable:

- **Camel Case** each word, except the first, starts with a capital letter: `myVariableName = "John"`

- **Pascal Case** each word starts with a capital letter: `MyVariableName = "John"`

- **Snake Case** each word is separated by an underscore character: `my_variable_name = "John"`
