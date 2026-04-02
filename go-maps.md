# GO MAPS

Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the `len()` function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

Go has multiple ways for creating maps.

- Create Maps Using var and :=
- Create Maps Using the make() Function

## Create an Empty Map

There are two ways to create an empty map. One is by using the make()function and the other is by using the following syntax.

## Allowed Key Types

The map key can be of any data type for which the equality operator (==) is defined. These include:

- Booleans
- Numbers
- Strings
- Arrays
- Pointers
- Structs
- Interfaces (as long as the dynamic type supports equality)

Invalid key types are:

- Slices
- Maps
- Functions

These types are invalid because the equality operator (==) is not defined for them.

## Allowed Value Types
The map values can be any type.

Removing elements from **Map** is done using the `delete()` function.

If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

You can use `range` to iterate over maps.

Maps are unordered data structures. If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order.

