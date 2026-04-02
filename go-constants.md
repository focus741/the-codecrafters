# GO CONSTANTS

---

If a variable should have a fixed value that cannot be changed, you can use `const` keyword. The `const` keyword declares the variable as "constant", which means that it is **unchangeable and read-only.**

```Syntax
const CONSTNAME type = value
```

---

## Declaring a Constant

Here is an example of declaring a constant in Go:

`const PI = 3.14`

## Constant Rules

Constant names follow the same naming rules as variables
Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
Constants can be declared both inside and outside of a function

--- 

## Constant Types

The two types of constants are:
**Typed constants** & **Untyped constants**

**Typed constants** are declared with a defined type: `const A int = 1
`

**Untyped Constants** are declaredwithout a type:
`const A = 1`

---

## Constants: Unchangeable and Read-only

When a constant is declared, it is not possible to change the value later: 
```
package main
import ("fmt")

func main() {
  const A = 1
  A = 2
  fmt.Println(A)
}
```

**Result:**
`./prog.go:8:7: cannot assign to A`

---

## Multiple Constants Declaration

Multiple constants can be grouped together into a block for readability:
```
package main
import ("fmt")

const (
  A int = 1
  B = 3.14
  C = "Hi!"
)

func main() {
  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
}
```
