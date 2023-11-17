# Binary Search Algorithm Challenge

Design a function to perform a binary search on a sorted array of integers to find the index of a given target integer. Binary search compares the target value to the middle element of the array; if they are not equal, the half in which the target cannot lie is eliminated, and the search continues on the remaining half until the target is found or the search space is empty.

**Function Description**

Write the `binarySearch` function with the following parameters:

-   _int arr[]:_ a sorted array of integers.
-   _int target:_ the integer to search for in the array.

**Returns**

-   _int:_ the index of the target integer in the array, or `-1` if the target is not found.

**Input Format**

Two arguments: a sorted array `arr` and the target integer `target`.

**Constraints**

-   1 <= len(arr) <= 10^6
-   -10^9 <= arr[i], target <= 10^9
-   `arr` is sorted in ascending order.

**Sample Input and Output**

1.  Input: arr = [1, 2, 3, 4, 5], target = 3 Output: 2

2.  Input: arr = [1, 5, 8, 12, 20], target = 5 Output: 1


**Edge Cases**

1.  Input: arr = [2], target = 2 Output: 0

    -   A single-element array should be correctly handled.
2.  Input: arr = [], target = 3 Output: -1

    -   An empty array should return `-1` since the target cannot be found.
3.  Input: arr = [1, 3, 5, 7], target = 6 Output: -1

    -   A target not present in the array should return `-1`.

**Explanation**

The binary search function should begin by determining the middle element of the array. If the middle element is equal to the target, the index of the middle element is returned. If the target is less than the middle element, the search continues on the left half of the array; if the target is greater, the search continues on the right half. The process repeats, narrowing down the potential locations of the target, until the target is found or until it is determined that the target is not in the array.
