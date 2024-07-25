package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Trim spaces and tabs from both ends
func trimSpaces(s string) string {
	start := 0
	end := len(s) - 1

	// Find the first non-space character
	for start <= end && (s[start] == ' ' || s[start] == '\t') {
		start++
	}

	// Find the last non-space character
	for end >= start && (s[end] == ' ' || s[end] == '\t') {
		end--
	}

	if start > end {
		return ""
	}

	return s[start : end+1]
}

// Split string into words
func splitWords(s string) []string {
	var words []string
	var word []rune

	for _, ch := range s {
		if ch == ' ' || ch == '\t' {
			if len(word) > 0 {
				words = append(words, string(word))
				word = []rune{}
			}
		} else {
			word = append(word, ch)
		}
	}

	if len(word) > 0 {
		words = append(words, string(word))
	}

	return words
}

// Join words with a single space
func joinWords(words []string) string {
	var result string
	for i, word := range words {
		if i > 0 {
			result += " "
		}
		result += word
	}
	return result
}

// Print string using z01 package
func printString(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func main() {
	if len(os.Args) != 2 {
		printString("\n")
		return
	}

	input := os.Args[1]
	trimmedInput := trimSpaces(input)
	if trimmedInput == "" {
		printString("\n")
		return
	}

	words := splitWords(trimmedInput)
	result := joinWords(words)
	printString(result + "\n")
}
