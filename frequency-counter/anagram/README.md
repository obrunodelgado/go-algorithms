
**Anagram Check Challenge**

Create a function that determines whether two strings are anagrams of each other. An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

The `areAnagrams` function should check if both strings contain the same letters in any order. The check should be case-insensitive and should not consider any non-letter characters. The function should count the occurrences of each letter in both strings and then compare these counts. If they match for all letters, the function returns `true`, indicating the strings are anagrams. If there is any discrepancy in the counts, it returns `false`.

**Function Description**

Construct the `areAnagrams` function with the following parameters:

-   _string word1:_ the first string.
-   _string word2:_ the second string.

**Returns**

-   _bool:_ `true` if the two strings are anagrams, `false` otherwise.

**Input Format**

Two string arguments, `word1` and `word2`, that will be checked to determine if they are anagrams.

**Constraints**

-   The strings will consist only of lowercase letters (a-z).
-   The length of both strings will be at least 1 character and at most 10^4 characters.

**Sample Input and Output**

1.  Input: "listen", "silent" Output: true

2.  Input: "hello", "billion" Output: false


**Edge Cases**

1.  Input: "finder", "Friend" Output: true

    -   The function should be case-insensitive.
2.  Input: "hello", "llheo" Output: true

    -   The order of characters is irrelevant; only the character counts matter.
3.  Input: "apple", "papel" Output: true

4.  Input: "rat", "cart" Output: false

    -   Strings of different lengths cannot be anagrams.

