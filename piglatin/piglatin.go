package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	word := os.Args[1]
	pigLatinWord := toPigLatin(word)
	printString(pigLatinWord)
}

func toPigLatin(word string) string {
	vowels := "aeiouAEIOU"
	vowelIndex := -1

	for i, char := range word {
		if containsRune(vowels, char) {
			vowelIndex = i
			break
		}
	}

	if vowelIndex == -1 {
		return "No vowels"
	}

	if vowelIndex == 0 {
		return word + "ay"
	}

	return word[vowelIndex:] + word[:vowelIndex] + "ay"
}

func containsRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

func printString(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
