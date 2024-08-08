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

	// Trim leading spaces and tabs
	start := 0
	for start < len(input) && (input[start] == ' ' || input[start] == '\t') {
		start++
	}

	// Trim trailing spaces and tabs
	end := len(input) - 1
	for end >= start && (input[end] == ' ' || input[end] == '\t') {
		end--
	}

	// If the resulting string is empty, return without printing anything
	if start > end {
		return
	}

	// Print words with exactly three spaces between them
	space := false
	wordCount := 0
	for i := start; i <= end; i++ {
		if input[i] == ' ' || input[i] == '\t' {
			space = true
		} else {
			if space && wordCount > 0 {
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				space = false
			}
			z01.PrintRune(rune(input[i]))
			if !space {
				wordCount++
			}
		}
	}
	z01.PrintRune('\n')
}
