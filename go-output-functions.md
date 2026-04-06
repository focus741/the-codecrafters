# GO OUTPUT FUNCTIONS

---

Go has three functions to output text: 

- Print()
- Println()
- Printf()

The **Print()** function prints its arguments with their default format.

The **Println** function is similar to **Print()** with the difference that a whitespace is added between the arguments, and a newline is added at the end.

The **Printf()** function first formats its argument based on the given formatting verb and then prints them.
Here are two formatting verbs; `%v` and `%T`.

**%v** is used to print the **value** of the arguments.
**%T** is used to print the **type** of the arguments.

---

# GO FORMATTING VERBS

Formatting verbs for Printf(): Go offers several formatting verbs that can be used with the **Printf()** function.

## General Formatting Verbs

The following verbs can be used with all data types:
**%v** prints the value in the default format.
**%#v** prints the value in Go-syntax format.
**%T** prints the type of the value.
**%%** prints the `%` sign.

---

## Integer Formatting Verbs

The following verbs can be used with the integer data type:

Verb:---------------------------------:Description

%b	Base 2

%d	Base 10

%+d	Base 10 and always show sign

%o	Base 8

%O	Base 8, with leading 0o

%x	Base 16, lowercase

%X	Base 16, uppercase

%#x	Base 16, with leading 0x

%4d	Pad with spaces (width 4, right justified)

%-4d	Pad with spaces (width 4, left justified)

%04d	Pad with zeroes (width 4
**Example**
```
package main
import ("fmt")
func main() {
  var i = 15
 
  fmt.Printf("%b\n", i)
  fmt.Printf("%d\n", i)
  fmt.Printf("%+d\n", i)
  fmt.Printf("%o\n", i)
  fmt.Printf("%O\n", i)
  fmt.Printf("%x\n", i)
  fmt.Printf("%X\n", i)
  fmt.Printf("%#x\n", i)
  fmt.Printf("%4d\n", i)
  fmt.Printf("%-4d\n", i)
  fmt.Printf("%04d\n", i)
}
```

**Result**
```
1111
15
+15
17
0o17
f
F
0xf
  15
15
0015
```

## String Formatting Verbs

The following verbs can be used with the string data type:
```
Verb	Description
%s	Prints the value as plain string
%q	Prints the value as a double-quoted string
%8s	Prints the value as plain string (width 8, right justified)
%-8s	Prints the value as plain string (width 8, left justified)
%x	Prints the value as hex dump of byte values
% x	Prints the value as hex dump with spaces
```
---

## Boolean Formatting Verbs

The following verb can be used with the boolean data type: 
`%t` value of the boolean operator in true or false format
(same as using %v)

**Example**
```
package main
import ("fmt")
func main() {
  var i = true
  var j = false
  fmt.Printf("%t\n", i)
  fmt.Printf("%t\n", j)
}
```
**Result**
```
true
false
```
## Float Formatting Verbs


The following verbs can be used with the float data type:

Verb	Description

%e	Scientific notation with 'e' as exponent

%f	Decimal point, no exponent

%.2f	Default width, precision 2

%6.2f	Width 6, precision 2

%g	Exponent as needed, only necessary digits

---
