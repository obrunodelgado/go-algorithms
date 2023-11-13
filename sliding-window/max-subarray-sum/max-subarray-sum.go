package main

import "fmt"

func main() {
	numeros1 := []int{1, 2, 5, 2, 8, 1, 5}
	fmt.Println(maxSubarraySum(numeros1, 2)) // 10
	numeros2 := []int{1, 2, 5, 2, 8, 1, 5}
	fmt.Println(maxSubarraySum(numeros2, 4)) // 17
	numeros3 := []int{4, 2, 1, 6}
	fmt.Println(maxSubarraySum(numeros3, 1)) // 6
	numeros4 := []int{4, 2, 1, 6, 2}
	fmt.Println(maxSubarraySum(numeros4, 4)) // 13
	var numeros5 []int
	fmt.Println(maxSubarraySum(numeros5, 4)) // 0
}

func maxSubarraySum(numbers []int, maxSequence int) int {
	if len(numbers) == 0 {
		return 0
	}

	sum := 0

	for i := 0; i < maxSequence; i++ {
		sum += numbers[i]
	}

	maxSum := sum

	for i := maxSequence; i < len(numbers); i++ {
		sum += numbers[i] - numbers[i-maxSequence]

		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
