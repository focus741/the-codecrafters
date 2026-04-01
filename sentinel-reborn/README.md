# THE INTERFACES

README.md must include:

- What the program does
- How to run it with examples
- Each transformation listed and explained
- What your personal contribution was
- One thing you found hardest today
- One thing you understand now that
  you did not understand this morning

# EJEME GODWIN - The Leader of this Magnificent Group

---

My Contribution today was epic. I tried my best actually even though i don't know much, i did what I know and it was tested and it ran well.

My work on this project was the capN which I did three times but i was cleared, i did the first one and they said it is an answer for a normal recoding question which I did another one and I never showed that to anybody bacause even me that wrote it was not convince, the code was working though but nah, so as asked for ideas I was shared some and I tried again and I got it right. By that time i was already tired and hungry.

My approach on the code I got right was that i use regexp for pattern matching, because it is like a search tool that can find, validate, or replace text based on rules instead of exact words. And I aslo used reCap to work with my text pattern and I also use ReplaceAllStringFunc to find all matches of a pattern and replace each one using a custom function and then the strings.Split which is used to to cut string into pieces using the separator.

Then the pushing and merge this is were the work is, I did the merging well but some file stress my life. Some are stubborn and refuse to merge, some of my group members work even refused to be pushed which almost killed me. lol

Lastly the main.go file where we join all the files together after the stress of pushing and merging. WE ARE DONE!!!

---

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



# Nnenna
I worked on the function that input "an" words starting with vowels and h.

* * i have to first convert the input to slice and then loop over it in order to find a standalone a, then check if the first letter of that word starts with vowel or h if it is so it will go back and convert the a to an.

* * one thing i found hardest today?

the one thing I found hardest today was to figure out how to get the standalone a and go forward to check the next word, then go back to convert

* * One thing you understand now that you did not understand this morning

Team work makes tedious job easier

---

# Quote and Punctuation Fixer by Agbaji Alexander

A robust Go utility designed to clean up messy text by normalizing whitespace, fixing incorrectly spaced quotes, and ensuring punctuation marks are properly attached to words.

Features
Whitespace Normalization: Automatically collapses multiple spaces into a single space and trims leading/trailing whitespace.

Smart Quote Handling:
Removes spaces immediately after an opening single quote.
Removes spaces immediately before a closing single quote.
Punctuation Correction: Detects periods (.) and ensures there is no trailing space between the preceding word and the period.
State Tracking: Uses a boolean toggle to distinguish between opening and closing quotes.

Logic Breakdown
Tokenization: The function uses strings.Fields and strings.Join to eliminate erratic spacing (e.g., "I am" becomes "I am").
String Builder: Uses strings.Builder for efficient string construction, minimizing memory allocations during the loop.

The inQuote Toggle:

First ': Marks the start of a quote. The logic checks the next character; if it's a space, it skips it to pull the quoted word tight to the mark.

Second ': Marks the end. It "looks back" at the strings.Builder buffer and removes any trailing space before placing the closing mark.
Punctuation Buffer: Similar to the closing quote logic, if a period is detected, the function trims any trailing space from the current result before adding the ..
Usage
input := " '  hello world  ' . "
output := fixQuotes(input)
// Result: "'hello world'."


Example Results
Input
Output
' awesome '
'awesome'
' I am the smartest guy . '
'I am the smartest guy.'
how they describe me: ' awesome '

---

# MR KENNEDY

---

Go Text Completer & Formatter
Project Overview
This project is a command-line tool written in Go that serves as an automated text editor. It transforms raw, messy input files into polished, grammatically correct text by applying a series of algorithmic modifications. The tool focuses on three primary areas: mathematical conversions, case transformations, and orthographic cleanup (punctuation and quotes).
Core Functionality: fixQuotes and formatText
The heart of this tool lies in its ability to handle erratic spacing. Unlike simple find-and-replace methods, the implementation uses a token-based approach and regular expressions to ensure precision.
1. Intelligent Quote Placement
The fixQuotes function uses a state-tracking boolean (inQuote) to handle pairs of single quotes.
Opening Quotes: It identifies the first instance of a quote and "pulls" the subsequent word toward it, eliminating any leading whitespace.
Closing Quotes: It identifies the second instance and "sucks" it back toward the preceding word, ensuring no space exists between the last character of the sentence and the closing mark.
2. Punctuation Normalization
The formatting logic ensures that standard punctuation marks (., ,, !, ?, :, ;) are:
Attached to the preceding word (no leading space).
Followed by exactly one space (unless at the end of the file).
Grouped correctly for special cases like ellipses (...) or combined marks (!?), preventing the logic from incorrectly spacing out these specific clusters.
Features
Number Conversions: Automatically detects (hex) and (bin) tags to convert preceding values into decimal format.
Case Modification: Supports (up), (low), and (cap) with optional numerical arguments to modify multiple preceding words.
Grammar Correction: Automatically converts the article "a" to "an" when the following word begins with a vowel or the letter 'h'.
Usage
bash
go run . input.txt output.txt


Use code with caution.
This command reads the contents of input.txt, applies all processing rules, and writes the finalized version to output.txt.
