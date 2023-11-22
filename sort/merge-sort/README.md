
# Merge Sort Algorithm Challenge

Craft a function that executes the merge sort algorithm to sort an array of integers. Merge sort is a divide-and-conquer algorithm that divides the input array into two halves, calls itself for the two halves, and then merges the two sorted halves.

**Function Description**

Implement the `mergeSort` function which sorts the array using the merge sort algorithm.

-   _int arr[]:_ an unsorted array of integers.

**Returns**

-   _int[]:_ the array sorted in ascending order.

**Input Format**

A single argument to the function, an array `arr` of integers that needs to be sorted.

**Constraints**

-   1 <= len(arr) <= 10^5
-   -10^9 <= arr[i] <= 10^9

**Sample Input and Output**

1.  Input: [38, 27, 43, 3, 9, 82, 10] Output: [3, 9, 10, 27, 38, 43, 82]

2.  Input: [5, 4, 3, 2, 1] Output: [1, 2, 3, 4, 5]


**Edge Cases**

1.  Input: [99] Output: [99]

    -   Arrays with a single element should return unchanged.
2.  Input: [] Output: []

    -   An empty array should return an empty array as there are no elements to sort.

**Explanation**

The function should recursively divide the array into halves until the base condition of a single element (or empty) array is reached. Then, it should start merging the arrays back together while sorting them. During the merge phase, the function should compare elements of the subarrays and combine them in sorted order. This process is repeated until the entire array is merged back together in a sorted manner.
