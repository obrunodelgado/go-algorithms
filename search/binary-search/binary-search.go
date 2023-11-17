package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	target := 3
	fmt.Println(binarySearch(numbers, target)) // 2

	numbers = []int{1, 5, 8, 12, 20}
	target = 5
	fmt.Println(binarySearch(numbers, target)) // 1

	numbers = []int{2}
	target = 2
	fmt.Println(binarySearch(numbers, target)) // -1

	var numbers2 []int
	target = 3
	fmt.Println(binarySearch(numbers2, target)) // -1

	numbers = []int{7, 3, 5, 7}
	target = 6
	fmt.Println(binarySearch(numbers, target)) // -1
}

func binarySearch(numbers []int, target int) int {
	left := 0
	right := len(numbers) - 1

	for left <= right {
		mid := left + (right-left)/2

		if numbers[mid] == target {
			return mid
		} else if numbers[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
