README.md must include:

- What the program does
- How to run it with examples
- Each transformation listed and explained
- What your personal contribution was
- One thing you found hardest today
- One thing you understand now that
  you did not understand this morning

# 
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

Lastly, i appended the remaining of the text after the marker to that slice and return it as a string through strings.Join.
---

# auoche

I build a function called capN which is capable of capitalising the first character of strings as asigned

What i did was i make sure that my program identifies the tag implimented in my text. 

knowing its start and end of a tag was my first logic. 

I also asign the word or words before the tag using a slice to be affected by the function of my tag using a loop. 

I emperward my tag with a simple .Title function with a finishing touch of joining the sliced words after convertion into strings as a ready processed words.


------------------------------------------------------------------------------------------
# Tavershima Elvis
**personal contribution;**
#A CLI tool for handling upper
. The CLI tool turns the previous word or words into upper case any where it sees (up, number) in a text the "(up, number)" is the command the Command Line Interface is using, if replace number in the bracket with a number in figure with the up, ;the texts at the length of the indicate number turns to uppercase.

 Example: input= the world of go(up, 4) is wild
          output= THE WORLD OF GO is wild
         
.My contribution to the program is the part where my function handles the ToUpper in the text and brainstorming with my group members in to bring out result in this aspect as our section was base on handling up, cap and low, also the general merging of the files.

**One Thing I Found Hardest and Understand Today**

.One thing i found hard today is the logic of implementing the command after identifying the markers in the text that represent commands

. one thing i understand is first, how to work in group, breaking task to peices and working on them also creating a pull request


# Onminyi Andrew Okala
**Personal Contribution**
I wrote the functions for `cap`, `low` and `up`
To run the code your input should carry this prefix `(cap)` for capitalization, `(low)` for lowercase and `(up)` for upppercase

Example:
fmt.Println(parseModifier("Onminyi Andrew Okala (up)"))
Output:
ONMINYI ANDREW OKALA

**My Difficult Moment Today**
At first i wasn't sure of how to implement the code, I even misplaced priority by going to perform the task assign to the other members of the group.

**What I understood Today**
I understood how to use the `strings.Index` and `strings.Contains()` in a function very well now.
---

# kinonoja

When we started; I was given the job of handling the conversion of hexadecimal to decimal. 
* HOW THIS WAS DONE
First, I had to use "strconv" because we were converting a string to an integer

Then i decided to declare a variable that err := strconv.ParseInt cause ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i.

After that i used if statement that if the error is not eqaul to it should return 0 and print "invalid hex string:" then i returned decimal and the nil figure

Then i had to wriite another function that will bring the input and return everything in the output figure as presented using regexp.Must.Compile

Then after i that, i joined in writing the read file and input to make sure the file brings out the output in everything
