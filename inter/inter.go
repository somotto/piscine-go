package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Check if there are exactly 2 command-line arguments
	if len(os.Args) != 3 {
		return // Display nothing
	}

	// Extract the two input strings
	str1, str2 := os.Args[1], os.Args[2]

	// Create a map to track characters from the first string
	seen := make(map[rune]bool)

	// Iterate over the first string and mark each character as unseen
	for _, char := range str1 {
		seen[char] = false
	}

	// Iterate over the second string and print common characters in order of str1
	for _, char := range str1 {
		if seen[char] == false {
			for _, char2 := range str2 {
				if char == char2 {
					z01.PrintRune(char)
					seen[char] = true // Mark as printed
					break
				}
			}
		}
	}

	z01.PrintRune('\n') // Add a newline
}
