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

	// Retrieve the input string from command-line argument
	input := os.Args[1]

	// Iterate over each character in the input string
	for _, char := range input {
		// Check if the character is a lowercase letter
		if char >= 'a' && char <= 'z' {
			// Convert lowercase letter to its mirror counterpart
			z01.PrintRune('z' - (char - 'a'))
		} else if char >= 'A' && char <= 'Z' {
			// Convert uppercase letter to its mirror counterpart
			z01.PrintRune('Z' - (char - 'A'))
		} else {
			// Non-alphabetical characters remain unchanged
			z01.PrintRune(char)
		}
	}

	// Print a newline at the end
	z01.PrintRune('\n')
}