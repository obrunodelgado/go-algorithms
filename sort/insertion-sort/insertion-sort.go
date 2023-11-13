package main

import "fmt"

func main() {
	numbers := []int{22, 27, 16, 2, 18, 6}
	fmt.Println(insertionSort(numbers)) // [2, 6, 16, 18, 22, 27]

	numbers = []int{-1, 6, 3, 2, 1, 9}
	fmt.Println(insertionSort(numbers)) // [-1, 1, 2, 3, 6, 9]

	numbers = []int{1}
	fmt.Println(insertionSort(numbers)) // 1

	var numbers2 []int
	fmt.Println(insertionSort(numbers2)) // []

	numbers = []int{4, 4, 4, 4}
	fmt.Println(insertionSort(numbers)) // [4, 4, 4, 4]
}

func insertionSort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {
		key := numbers[i]
		j := i - 1

		for j >= 0 && numbers[j] > key {
			numbers[j+1] = numbers[j]
			j--
		}

		numbers[j+1] = key
	}

	return numbers
}
