package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for _, s := range args {
		var result string

		for i, ch := range s {
			if i != len(s)-1 && s[i+1] == ' ' {
				result += string(ToUpper(ch))
			} else if i == len(s)-1 {
				result += string(ToUpper(ch))
			} else {
				result += string(ToLower(ch))
			}
		}
		for _, char := range result {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

func ToUpper(n rune) rune {
	if n >= 'a' && n <= 'z' {
		return n - 32
	}
	return n
}

func ToLower(n rune) rune {
	if n >= 'A' && n <= 'Z' {
		return n + 32
	}
	return n
}