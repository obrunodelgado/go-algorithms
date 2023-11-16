package main

import "fmt"

func main() {
	numbers := []int{29, 10, 14, 37, 13}
	fmt.Println(selectionSort(numbers)) // [10, 13, 14, 29, 37]

	numbers = []int{32, 3, 2, 10}
	fmt.Println(selectionSort(numbers)) // [2, 3, 10, 32]

	numbers = []int{3}
	fmt.Println(selectionSort(numbers)) // 3

	var numbers2 []int
	fmt.Println(selectionSort(numbers2)) // []

	numbers = []int{5, -2, -5, 2}
	fmt.Println(selectionSort(numbers)) // [-5, -2, 2, 5]

	numbers = []int{1, 1, 1, 1}
	fmt.Println(selectionSort(numbers)) // [1, 1, 1, 1]
}

func selectionSort(numbers []int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		min := i

		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[min] {
				min = j
			}
		}

		if min > i {
			aux := numbers[i]
			numbers[i] = numbers[min]
			numbers[min] = aux
		}
	}

	return numbers
}
