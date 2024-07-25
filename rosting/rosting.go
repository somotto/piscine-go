package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Split the string into words
func split(s string) []string {
	var words []string
	word := ""

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word += string(s[i])
		} else {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		}
	}
	if word != "" {
		words = append(words, word)
	}

	return words
}

// Join words with a single space
func join(words []string) string {
	if len(words) == 0 {
		return ""
	}

	result := words[0]
	for i := 1; i < len(words); i++ {
		result += " " + words[i]
	}

	return result
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	input := os.Args[1]
	words := split(input)

	if len(words) == 0 {
		z01.PrintRune('\n')
		return
	}

	rotated := append(words[1:], words[0])
	output := join(rotated)

	for _, r := range output {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
