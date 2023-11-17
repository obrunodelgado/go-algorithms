package main

import "fmt"

func main() {
	word1 := "listen"
	word2 := "silent"
	fmt.Println(areAnagrams(word1, word2)) // true

	word1 = "listena"
	word2 = "silent"
	fmt.Println(areAnagrams(word1, word2)) // false

	word1 = "hello"
	word2 = "billion"
	fmt.Println(areAnagrams(word1, word2)) // false

	word1 = "apple"
	word2 = "papel"
	fmt.Println(areAnagrams(word1, word2)) // true

	word1 = "rat"
	word2 = "cart"
	fmt.Println(areAnagrams(word1, word2)) // false
}

func areAnagrams(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	frequency := make(map[rune]int)
	for _, value := range word1 {
		frequency[value] += 1
	}

	for _, value := range word2 {
		if _, found := frequency[value]; found == false {
			return false
		} else {
			frequency[value] -= 1
		}
	}

	return true
}
