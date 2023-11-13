
**Duplicate Letters Count Problem**

Design a function that accepts a string and calculates the number of letters that appear more than once in the given string.

**Example**

`word = "letters"`

The function must count the number of unique letters that repeat in the word "letters". There are two letters that appear more than once: 'e' appears twice, and 't' also appears twice.

**Function Description**

Write the function `countDuplicateLetters`.

`countDuplicateLetters` has the following parameter(s):

-   _string word:_ the word to be analyzed

**Returns**

-   _int:_ the number of unique letters that repeat

**Input Format** A single string argument representing the word to be analyzed.

**Constraints**

-   The string will contain only lowercase letters (a-z).
-   The length of the string `word` will be at least 1 and at most 10^4 characters.

**Sample Input and Output**

    Input: "love"
    Output: 0
    
    Input: "letters"
    Output: 2
    
    Input: "repeats"
    Output: 1

**Explanation**

In the input "love", there are no letters that repeat, hence the output is 0. For "letters", the letter 'e' appears twice and the letter 't' also appears twice, thus the output is 2 with the repeating letters listed. In "repeats", the letter 'e' is the only one that repeats, resulting in the output 1, with 'e' indicated as the repeating letter.
