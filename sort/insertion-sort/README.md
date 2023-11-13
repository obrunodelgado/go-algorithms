Insertion Sort Algorithm Challenge

Develop a function that utilizes the insertion sort technique to order an array of integers. Insertion sort is an algorithm that builds the final sorted array one item at a time. It is much less efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort, but has the advantage of being simple to understand and can perform well on small lists.

Function Description

Implement the insertionSort function with the following specification:

    int arr[]: an array of integers that needs to be sorted.

Returns

    int[]: the array sorted in ascending order.

Input Format

A single argument, an integer array arr that is to be sorted using insertion sort.

Constraints

    1 <= len(arr) <= 10^3
    -10^3 <= arr[i] <= 10^3

Sample Input and Output

    Input: [22, 27, 16, 2, 18, 6]
    Output: [2, 6, 16, 18, 22, 27]

    Input: [-1, 6, 3, 2, 1, 9]
    Output: [-1, 1, 2, 3, 6, 9]

Edge Cases

    Input: [1]
    Output: [1]
        Arrays with a single element should return the element as the sorted array.

    Input: []
    Output: []
        An empty array should return an empty array.

    Input: [4, 4, 4, 4]
    Output: [4, 4, 4, 4]
        Arrays with identical elements should return the same array.

Explanation

Insertion sort works by taking elements from the unsorted array and inserting them at the correct position in a new sorted array. This is done by moving elements greater than the new element one position up to make space for the inserted element. It is an in-place, stable sorting algorithm, which is easy to implement, and is efficient for small datasets.