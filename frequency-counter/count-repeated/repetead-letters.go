package main

import "fmt"

func main() {
	word := "love"
	fmt.Println(countRepeatedLetters(word))

	word = "letters"
	fmt.Println(countRepeatedLetters(word))

	word = "repeats"
	fmt.Println(countRepeatedLetters(word))

	word = "aaabbbcc"
	fmt.Println(countRepeatedLetters(word))
}

func countRepeatedLetters(word string) int {
	frequency := make(map[rune]int)

	for _, value := range word {
		frequency[value] += 1
	}

	counter := 0
	for _, value := range frequency {
		if value > 1 {
			counter++
		}
	}

	return counter
}
