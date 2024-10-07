package main

import (
	"fmt"
	"unicode"
)

func countLetters(str string) map[rune]int {
	letterCount := make(map[rune]int)

	for _, char := range str {
		// Check if the character is a letter
		if unicode.IsLetter(char) {
			letterCount[char]++
		}
	}

	return letterCount
}

func main() {
	str := "Hello, Go! How are you?"

	// Count letters in the string
	counts := countLetters(str)

	// Print the letter counts
	for letter, count := range counts {
		fmt.Printf("%c: %d\n", letter, count)
	}
}
