package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		z01.PrintRune('\n')
		return
	}

	str1 := args[0]
	str2 := args[1]

	seen := make(map[rune]bool)

	for _, char := range str1 + str2 {
		if !seen[char] {
			z01.PrintRune(char)
			seen[char] = true
		}
	}

	z01.PrintRune('\n')
}
