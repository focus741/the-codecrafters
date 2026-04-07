# GO SLICES

Slices are similar to arrays, but are more powerful and flexible. Like arrays , slices are also used to store multiple values of the same type in a single variable. Unlike arrays, the length of a slice can grow and shrink as you see fi.

In Go, there are several ways to create a slice:

- Using the []datatype{values} format
- Create a slice from an array
- Using the make() function

In Go, there are two functions that can be used to return the length and capacity of a slice:

- `len()` function - returns the length of the slice (the number of elements in the slice)
- `cap()` function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

---

## Go Access, Change, Append and Copy Slices

### Access Elements of a Slice

You can access a specific slice element by referring to index number.

In Go, indexes start at 0, that means that [0] is the first element , [1] is the second element, etc.

### Change Elements of a Slice

You can also change a specific slice element by referring to the index number.



### Append Elements To a Slice

You can append elements to the end of a slice using the append()function

### Append One Slice To Another Slice

To append all the elements of one slice to another slice, use the append()function

### Change The Length of a Slice

Unlike arrays, it is possible to change the length of a slice.
