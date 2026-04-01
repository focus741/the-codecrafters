## hmusa

I implemented the function `lowN()` that handles the transformation of a specified number of words in a string to lowercase

How this was done:
```
First, i found the marker by identifying the index of the opening and the closing parentheses

Secondly, i put the words before the marker into a slice so they'll be easier to manipulate through indexing

Thirdly, i looped inside the marker to find the number, then created another loop starting at the lenght of the slice minus that number, then transform everything from that point onwards

Lastly, i appended the remaining of the text after the marker to that slice and return it as a string through strings.Join.
```
---