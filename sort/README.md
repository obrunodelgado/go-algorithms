
**Bubble Sort Algorithm Problem**

Implement a function that applies the bubble sort technique, effectively sorting an array of integers in non-decreasing order. The bubble sort algorithm repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order. The pass through the list is repeated until the list is sorted.

**Function Description**

Write the `bubbleSort` function with the following parameter:

-   _int arr[]:_ an array of integers to be sorted.

**Returns**

-   _int[]:_ the sorted array in non-decreasing order.

**Input Format**

A single argument to the function, an array `arr` of integers that needs to be sorted.

**Constraints**

-   1 <= len(arr) <= 10^3
-   -10^3 <= arr[i] <= 10^3

**Sample Input and Output**

1.  Input: [5, 1, 4, 2, 8] Output: [1, 2, 4, 5, 8]

2.  Input: [32, 3, 2, 10] Output: [2, 3, 10, 32]


**Edge Cases**

1.  Input: [3] Output: [3]

    -   Single-element arrays should return the same single element.
2.  Input: [] Output: []

    -   An empty array should return an empty array since there are no elements to sort.
3.  Input: [5, -2, -5, 2] Output: [-5, -2, 2, 5]

    -   The function should be able to handle negative numbers as well.
4.  Input: [1, 1, 1, 1] Output: [1, 1, 1, 1]

    -   Arrays where all elements are the same should return the identical array.

**Explanation**

The function performs the bubble sort algorithm, which involves nested loops: an outer loop to control the number of passes, and an inner loop to perform individual comparisons between elements. With each complete pass, the largest unsorted element bubbles up to its correct position at the end of the array. This process repeats until no further swaps are necessary, indicating that the array is sorted.
