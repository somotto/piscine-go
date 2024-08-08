package main

import (
	"os"

	"github.com/01-edu/z01"
)

// CapitalizeLastLetter capitalizes the last letter of each word in the input string
func CapitalizeLastLetter(s string) string {
	result := ""
	word := ""
	for i, r := range s {
		if r == ' ' {
			if len(word) > 0 {
				// Convert the whole word to lowercase
				word = lowerCase(word)
				// Capitalize the last letter
				lastLetter := runeToUpper(rune(word[len(word)-1]))
				word = word[:len(word)-1] + string(lastLetter)
				result += word
				word = ""
			}
			result += " "
		} else {
			word += string(r)
		}
		// Process the last word in the string
		if i == len(s)-1 && len(word) > 0 {
			word = lowerCase(word)
			lastLetter := runeToUpper(rune(word[len(word)-1]))
			word = word[:len(word)-1] + string(lastLetter)
			result += word
		}
	}
	return result
}

// lowerCase converts a string to lowercase
func lowerCase(s string) string {
	result := ""
	for _, r := range s {
		result += string(runeToLower(r))
	}
	return result
}

// runeToUpper converts a rune to uppercase if it's a lowercase letter
func runeToUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}

// runeToLower converts a rune to lowercase if it's an uppercase letter
func runeToLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

// PrintFormattedArgs processes the command line arguments and prints them
func PrintFormattedArgs() {
	args := os.Args[1:] // Get command line arguments
	if len(args) == 0 {
		return
	}

	for i, arg := range args {
		formatted := CapitalizeLastLetter(arg)
		for _, r := range formatted {
			z01.PrintRune(r)
		}
		if i < len(args)-1 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintFormattedArgs()
}
