package main

import (
	"os"

	"github.com/01-edu/z01"
)

// reverseWords reverses the words in the given string
func reverseWords(s string) string {
	// Split the string into words
	words := []string{}
	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			words = append(words, s[start:i])
			start = i + 1
		}
	}

	// Reverse the words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Join the words with a single space and return the result
	result := ""
	for i, word := range words {
		if i > 0 {
			result += " "
		}
		result += word
	}
	return result
}

func main() {
	// Check if the number of arguments is exactly 2 (program name and one input string)
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]

	// Reverse the words in the input string
	reversed := reverseWords(input)

	// Print the reversed string followed by a newline
	for _, char := range reversed {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
