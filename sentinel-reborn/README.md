README.md must include:

- What the program does
- How to run it with examples
- Each transformation listed and explained
- What your personal contribution was
- One thing you found hardest today
- One thing you understand now that
  you did not understand this morning


# hmusa

**What your personal contribution was**

I implemented the function `lowN()` that handles the transformation of a specified number of words in a string to lowercase

How this was done:
First, i found the marker by identifying the index of the opening and the closing parentheses

Secondly, i put the words before the marker into a slice so they'll be easier to manipulate through indexing

Thirdly, i looped inside the marker to find the number, then created another loop starting at the lenght of the slice minus that number, then transform everything from that point onwards

Lastly, i appended the remaining of the text after the marker to that slice and return it as a string through strings.Join.

**One thing you found hardest today**

What i found hardest was using the number gotten from the marker to transform that number of words in the slice from the end of the slice.

**One thing you understand now that you did not understand this morning**

I learnt new tactics to work around my loops, slices and strings, my creativity is what keeps me open minded to solving solutions i didn't have answers to.

---


## auoche

I build a function called capN which is capable of capitalising the first character of strings as asigned

What i did was i make sure that my program identifies the tag implimented in my text. knowing its start and end of a tag was my first logic then i also asign the word or words before the tag using a slice to be affected by the function of my tag using a loop then i emperward my tag with a simple .Title function with a finishing touch of joining the sliced words after convertion into strings as a ready processed words.
------------------------------------------------------------------------------------------
# Tavershima Elvis
# personal contribution;
#A CLI tool for handling upper
. The CLI tool turns the previous word or words into upper case any where it sees (up, number) in a text the "(up, number)" is the command the Command Line Interface is using, if replace number in the bracket with a number in figure with the up, ;the texts at the length of the indicate number turns to uppercase.

 Example: input= the world of go(up, 4) is wild
          output= THE WORLD OF GO is wild
         
.My contribution to the program is the part where my function handles the ToUpper in the text and brainstorming with my group members in to bring out result in this aspect as our section was base on handling up, cap and low, also the general merging of the files.

# One Thing I Found Hardest and Understand Today

.One thing i found hard today is the logic of implementing the command after identifying the markers in the text that represent commands

. one thing i understand is first, how to work in group, breaking task to peices and working on them also creating a pull request

