package main

import "fmt"

// FirstWord returns the first word of the string followed by a newline
func FirstWord(s string) string {
	// Trim leading spaces and tabs
	start := 0
	for start < len(s) && (s[start] == ' ' || s[start] == '\t') {
		start++
	}

	// If the string is empty or only contains spaces/tabs, return a newline
	if start == len(s) {
		return "\n"
	}

	// Find the end of the first word
	end := start
	for end < len(s) && s[end] != ' ' && s[end] != '\t' {
		end++
	}

	// Extract the first word
	firstWord := s[start:end]

	// Return the first word followed by a newline
	return firstWord + "\n"
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("      hello   .........  bye"))
}
