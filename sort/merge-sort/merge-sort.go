package main

import "fmt"

func main() {
	numbers := []int{10, 24, 76, 73, 72, 1, 9} // 1, 9, 10, 24, 72, 73, 76
	fmt.Println(mergeSort(numbers))

	numbers = []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println(mergeSort(numbers)) // 3, 9, 10, 27, 38, 43, 82

	numbers = []int{5, 4, 3, 2, 1}
	fmt.Println(mergeSort(numbers)) // 1, 2, 3, 4, 5

	numbers = []int{99}
	fmt.Println(mergeSort(numbers)) // 99

	var numbers1 []int
	fmt.Println(mergeSort(numbers1)) // []
}

func mergeSort(numbers []int) []int {
	size := len(numbers)

	if size <= 1 {
		return numbers
	}

	mid := len(numbers) / 2
	left := mergeSort(numbers[:mid])
	right := mergeSort(numbers[mid:])

	return merge(left, right)
}

func merge(first []int, second []int) []int {
	i := 0
	j := 0
	var result []int

	for i < len(first) && j < len(second) {
		if first[i] > second[j] {
			result = append(result, second[j])
			j++
		} else {
			result = append(result, first[i])
			i++
		}
	}

	for i < len(first) {
		result = append(result, first[i])
		i++
	}

	for j < len(second) {
		result = append(result, second[j])
		j++
	}

	return result
}
