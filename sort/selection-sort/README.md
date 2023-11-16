

# Selection Sort Algorithm Challenge


Develop a function that implements the selection sort algorithm, which sorts an array by repeatedly finding the minimum element from the unsorted part and moving it to the beginning.


## Function Description


Implement the selectionSort function with the following parameters:

int arr[]: an unsorted array of integers.  
Returns

int[]: the array sorted in ascending order.  
Input Format

A single argument, an array arr of integers that is to be sorted using selection sort.


## Constraints


1 <= len(arr) <= 10^3 -10^3 <= arr[i] <= 10^3

## Sample Input and Output

1. Input: [29, 10, 14, 37, 13]
   Output: [10, 13, 14, 29, 37]
2. Input: [5, 3, 8, 4, 2]
   Output: [2, 3, 4, 5, 8]

**Edge Cases**

2. Input: [1]
   Output: [1]
   Single-element arrays should remain unchanged.
3. Input: []
   Output: []
   An empty array should return an empty array.

**Explanation**  
Selection sort works by dividing the input list into two parts: a sorted sublist of items which is built up from left to right at the front (left) of the list, and a sublist of the remaining unsorted items that occupy the rest of the list. Initially, the sorted sublist is empty and the unsorted sublist is the entire input list. The algorithm proceeds by finding the smallest (or largest, depending on sorting order) element in the unsorted sublist, exchanging it with the leftmost unsorted element (putting it in sorted order), and moving the sublist boundaries one element to the right.
