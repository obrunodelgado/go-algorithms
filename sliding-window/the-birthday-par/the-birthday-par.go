package main

import "fmt"

func main() {
	valores := []int{1, 2, 1, 3, 2}
	fmt.Println(countPossibilities(valores, 3, 2))

	valores = []int{1, 1, 1, 1, 1, 1}
	fmt.Println(countPossibilities(valores, 3, 2))

	valores = []int{4}
	fmt.Println(countPossibilities(valores, 4, 1))
}

func countPossibilities(numbers []int, day int, month int) int {
	maxSum := 0
	possibilities := 0

	for i := 0; i < month; i++ {
		maxSum += numbers[i]
	}

	if maxSum == day {
		possibilities++
	}

	for i := month; i < len(numbers); i++ {
		maxSum += numbers[i] - numbers[i-1]

		if maxSum == day {
			possibilities++
		}
	}

	return possibilities
}
