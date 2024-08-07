package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return // Exit if the number of arguments is not 1
	}

	input := os.Args[1]
	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			// Transform lowercase letters
			z01.PrintRune((char-'a'+13)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			// Transform uppercase letters
			z01.PrintRune((char-'A'+13)%26 + 'A')
		} else {
			// Print non-letter characters unchanged
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n') // Print newline
}
