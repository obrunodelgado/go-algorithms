
Two children, Lily and Ron, want to share a chocolate bar. Each of the squares has an integer on it.

Lily decides to share a contiguous segment of the bar selected such that:

-   The length of the segment matches Ron's birth month, and,
-   The sum of the integers on the squares is equal to his birth day.

Determine how many ways she can divide the chocolate.

**Example**

    s = [2, 2, 1, 3, 2]
    d = 4
    m = 2



Lily wants to find segments summing to Ron's birth day, `d = 4` with a length equalling his birth month, `m = 2`. In this case, there are two segments meeting her criteria: `[2, 2]` and `[1, 3]`.

**Function Description**

Complete the _birthday_ function in the editor below.

birthday has the following parameter(s):

-   _int s[n]:_ the numbers on each of the squares of chocolate
-   _int d:_ Ron's birth day
-   _int m:_ Ron's birth month

**Returns**

-   _int:_ the number of ways the bar can be divided

**Input Format**
The first argument contains integers `s[i]` representing the numbers on the chocolate squares where `0 <= i < n`.  
The second argument is an integer containing Ron's birth day
The thirdy argument is an integer containing Ron's birth month

**Constraints**

- 1 <= `n` <= 100
- 1 <= `s[i]` <= 5, where (0 <= `i` < `n`)
- 1 <= `d` <= 31
- 1 <= `m` <= 12

**Sample Input**

    5
    1 2 1 3 2
    3 2

**Sample Output**

    2

For more details, visit https://www.hackerrank.com/challenges/the-birthday-bar/problem
