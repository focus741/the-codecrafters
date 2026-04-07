# GO STRUCT

## Go Structures

A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

A struct can be useful for grouping data together to create records.

## Declare a Struct

To declare a structure in Go, use the `type` and `struct` keywords.

## Pass Struct as Function Arguments
You can also pass a structure as a function argument, like this:

**Example**
```
package main
import ("fmt")
type Person struct {
  name string
  age int
  job string
  salary int
}
func main() {
  var pers1 Person
  var pers2 Person
  // Pers1 specification
  pers1.name = "Hege"
  pers1.age = 45
  pers1.job = "Teacher"
  pers1.salary = 6000
  // Pers2 specification
  pers2.name = "Cecilie"
  pers2.age = 24
  pers2.job = "Marketing"
  pers2.salary = 4500
  // Print Pers1 info by calling a function
  printPerson(pers1)
  // Print Pers2 info by calling a function
  printPerson(pers2)
}
func printPerson(pers Person) {
  fmt.Println("Name: ", pers.name)
  fmt.Println("Age: ", pers.age)
  fmt.Println("Job: ", pers.job)
  fmt.Println("Salary: ", pers.salary)
}
```

**Result:**

```
Name: Hege
Age: 45
Job: Teacher
Salary: 6000
Name: Cecilie
Age: 24
Job: Marketing
Salary: 4500
```
