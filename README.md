# Go Text Modifier Tool

This is a small Go program I made to edit text files automatically. It can do things like convert numbers, change word cases, fix punctuation, and even correct “a” to “an” when needed.

## What it can do

- Convert hexadecimal `(hex)` and binary `(bin)` numbers to decimal  
- Change words to uppercase `(up)`, lowercase `(low)`, or capitalize `(cap)`  
- Works on a single word or a number of words if you specify `(up, 2)`  
- Fix punctuation so it sits correctly next to words  
- Adjust single quotes `' '` around words or phrases  
- Switch “a” to “an” when the next word starts with a vowel or “h”  

## How to use it

Just run the program with your input file and the output file name:

```bash
$ go run . input.txt output.txt
Input (sample.txt):

it (cap) was the best of times, it was the worst of times (up)


Output (result.txt):

It was the best of times, it was the worst of TIMES
