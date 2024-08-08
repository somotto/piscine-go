package main

import "fmt"

// LastWord returns the last word of the string followed by a newline
func LastWord(s string) string {
	// Trim trailing spaces and tabs
	end := len(s) - 1
	for end >= 0 && (s[end] == ' ' || s[end] == '\t') {
		end--
	}

	// If the string is empty or only contains spaces/tabs, return a newline
	if end < 0 {
		return "\n"
	}

	// Find the start of the last word
	start := end
	for start >= 0 && s[start] != ' ' && s[start] != '\t' {
		start--
	}

	// Extract the last word
	lastWord := s[start+1 : end+1]

	// Return the last word followed by a newline
	return lastWord + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
