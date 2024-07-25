package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	words := split(input, ' ')
	reverse(words)
	printWords(words)
}

func split(s string, sep rune) []string {
	var words []string
	start := 0
	for i, c := range s {
		if c == sep {
			words = append(words, s[start:i])
			start = i + 1
		}
	}
	words = append(words, s[start:])
	return words
}

func reverse(words []string) {
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
}

func printWords(words []string) {
	for i, word := range words {
		for _, c := range word {
			z01.PrintRune(c)
		}
		if i < len(words)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
