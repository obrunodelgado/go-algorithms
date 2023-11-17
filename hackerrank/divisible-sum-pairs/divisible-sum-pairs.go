package main

import "fmt"

func main() {
	ar := []int{1, 2, 3, 4, 5, 6}
	k := 5
	fmt.Println(divisibleSumPairs(ar, k)) // 3

	ar = []int{1, 3, 2, 6, 1, 2}
	k = 3
	fmt.Println(divisibleSumPairs(ar, k)) // 5
}

func divisibleSumPairs(ar []int, k int) int {
	frequency := make(map[int]int)
	for _, value := range ar {
		remainder := value % k
		frequency[remainder]++
	}

	count := 0

	for i := 0; i <= k/2; i++ {
		if i == 0 || (k%2 == 0 && i == k/2) {
			count += (frequency[i]*frequency[i] - 1) / 2
		} else {
			count += frequency[i] * frequency[k-1]
		}
	}

	return count
}
