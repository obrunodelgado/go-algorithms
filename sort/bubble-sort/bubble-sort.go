package main

import "fmt"

func main() {
	numbers := []int{5, 1, 4, 2, 8}
	fmt.Println(bubbleSort(numbers)) // [1, 2, 4, 5, 8]

	numbers = []int{32, 3, 2, 10}
	fmt.Println(bubbleSort(numbers)) // [2, 3, 10, 32]

	numbers = []int{3}
	fmt.Println(bubbleSort(numbers)) // 3

	var numbers2 []int
	fmt.Println(bubbleSort(numbers2)) // []

	numbers = []int{5, -2, -5, 2}
	fmt.Println(bubbleSort(numbers)) // [-5, -2, 2, 5]

	numbers = []int{1, 1, 1, 1}
	fmt.Println(bubbleSort(numbers)) // [1, 1, 1, 1]
}

func bubbleSort(numbers []int) []int {
	higherValuePosition := len(numbers) - 1
	var swapped bool

	for {
		swapped = false

		for i := 0; i < higherValuePosition; i++ {
			if numbers[i] > numbers[i+1] {
				aux := numbers[i]
				numbers[i] = numbers[i+1]
				numbers[i+1] = aux
				swapped = true
			}
		}

		higherValuePosition--

		if !swapped {
			break
		}
	}
	return numbers
}
