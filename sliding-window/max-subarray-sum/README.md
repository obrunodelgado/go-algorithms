
**Max Sum Subarray Problem**

Given an array of integers, your task is to write a function that takes this array and a number `n` as arguments. The function should calculate the maximum sum of `n` consecutive elements in the array.

**Example**

`s = [1, 2, 5, 2, 8, 1, 5]` `n = 4`

The function should find the maximum sum of `n = 4` consecutive elements in the array `s`. The maximum sum in this case is from the subarray `[2, 5, 2, 8]`, which adds up to `17`.

**Function Description**

Complete the `maxSumSubarray` function in the editor below.

`maxSumSubarray` has the following parameter(s):

-   _int s[]:_ the array of integers
-   _int n:_ the number of consecutive elements to sum

**Returns**

-   _int:_ the maximum sum of `n` consecutive elements

**Input Format** The first argument is an array of integers representing the numbers in the array. The second argument is an integer `n` representing the number of consecutive elements to sum.

**Constraints**

-   1 <= `len(s)` <= 10^5
-   -10^4 <= `s[i]` <= 10^4, where (0 <= `i` < `len(s)`)
-   1 <= `n` <= `len(s)`

**Sample Input and Output**

    Input: [1, 2, 5, 2, 8, 1, 5], n = 4
    Output: 17
    
    Input: [4, 2, 1, 6], n = 1
    Output: 6
    
    Input: [4, 2, 1, 6, 2], n = 4
    Output: 13

**Explanation**

For the first sample, the maximum sum of 4 consecutive elements in the array `[1, 2, 5, 2, 8, 1, 5]` is 17, achieved by summing up `[2, 5, 2, 8]`. For the second sample, the maximum sum of 1 element in the array `[4, 2, 1, 6]` is 6, as the highest single number is 6. For the third sample, the maximum sum of 4 consecutive elements in the array `[4, 2, 1, 6, 2]` is 13, achieved by summing up `[4, 2, 1, 6]`.
