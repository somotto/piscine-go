package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	input := os.Args[1]

	// Remove leading spaces and tabs
	start := 0
	for start < len(input) && (input[start] == ' ' || input[start] == '\t') {
		start++
	}

	// Remove trailing spaces and tabs
	end := len(input) - 1
	for end >= start && (input[end] == ' ' || input[end] == '\t') {
		end--
	}

	// If the resulting string is empty, print a newline and return
	if start > end {
		z01.PrintRune('\n')
		return
	}

	// Print words with a single space between them
	space := false
	for i := start; i <= end; i++ {
		if input[i] == ' ' || input[i] == '\t' {
			space = true
		} else {
			if space {
				z01.PrintRune(' ')
				space = false
			}
			z01.PrintRune(rune(input[i]))
		}
	}
	z01.PrintRune('\n')
}
