# Quick Sort Algorithm Challenge

Develop a function that implements the quick sort algorithm for sorting an array of integers. Quick sort is a highly efficient sorting algorithm that uses a divide-and-conquer approach. It selects an element as a pivot and partitions the given array around the chosen pivot.

**Function Description**

Create the `quickSort` function to sort the array using the quick sort method.

-   _int arr[]:_ an array of integers to be sorted.

**Returns**

-   _int[]:_ the array sorted in ascending order.

**Input Format**

A single argument, an array `arr` of integers that needs to be sorted using quick sort.

**Constraints**

-   1 <= len(arr) <= 10^5
-   -10^9 <= arr[i] <= 10^9

**Sample Input and Output**

1.  Input: [10, 7, 8, 9, 1, 5] Output: [1, 5, 7, 8, 9, 10]

2.  Input: [6, 3, 9, 5, 2, 8] Output: [2, 3, 5, 6, 8, 9]


**Edge Cases**

1.  Input: [2] Output: [2]

    -   Arrays with a single element should return unchanged.
2.  Input: [] Output: []

    -   An empty array should return an empty array.
3.  Input: [3, 3, 3] Output: [3, 3, 3]

    -   Arrays with identical elements should be handled correctly.

**Explanation**

The function should choose a pivot element from the array and partition the other elements into two sub-arrays, according to whether they are less than or greater than the pivot. The sub-arrays are then sorted recursively. This process continues until the base case of an array with zero or one element is reached, which are already sorted. The end result is a fully sorted array. The choice of pivot and the partitioning scheme can vary, but a common approach is to use the last element as the pivot and then to move elements smaller than the pivot to the left and larger ones to the right.
