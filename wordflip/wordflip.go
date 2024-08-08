package main

import "fmt"

// WordFlip reverses the order of words in a string
func WordFlip(str string) string {
	// Step 1: Trim leading and trailing spaces
	start := 0
	end := len(str) - 1

	// Find the start index of the first non-space character
	for start <= end && (str[start] == ' ' || str[start] == '\t') {
		start++
	}

	// Find the end index of the last non-space character
	for end >= start && (str[end] == ' ' || str[end] == '\t') {
		end--
	}

	// If the resulting string is empty
	if start > end {
		return "Invalid Output\n"
	}

	// Step 2: Normalize multiple spaces between words
	trimmedStr := ""
	inWord := false
	for i := start; i <= end; i++ {
		if str[i] == ' ' || str[i] == '\t' {
			if inWord {
				trimmedStr += " "
				inWord = false
			}
		} else {
			trimmedStr += string(str[i])
			inWord = true
		}
	}

	// Step 3: Reverse the order of words
	words := []string{}
	start = 0
	for i := 0; i <= len(trimmedStr); i++ {
		if i == len(trimmedStr) || trimmedStr[i] == ' ' {
			if start < i {
				words = append(words, trimmedStr[start:i])
			}
			start = i + 1
		}
	}

	reversedStr := ""
	for i := len(words) - 1; i >= 0; i-- {
		if i < len(words)-1 {
			reversedStr += " "
		}
		reversedStr += words[i]
	}

	return reversedStr + "\n"
}

func main() {
	// Example test cases
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
