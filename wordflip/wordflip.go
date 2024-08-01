package main

import (
	"fmt"
)

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}

func WordFlip(str string) string {
	if len(str) == 0 {
		return "Invalid Output\n"
	}

	// Trim leading and trailing spaces
	start, end := 0, len(str)-1
	for start < len(str) && str[start] == ' ' {
		start++
	}
	for end >= 0 && str[end] == ' ' {
		end--
	}
	if start > end {
		return "\n"
	}

	// Split into words and reverse
	words := []string{}
	word := ""
	for i := start; i <= end; i++ {
		if str[i] != ' ' {
			word += string(str[i])
		} else if word != "" {
			words = append([]string{word}, words...)
			word = ""
		}
	}
	if word != "" {
		words = append([]string{word}, words...)
	}

	// Join words
	result := ""
	for i, w := range words {
		if i > 0 {
			result += " "
		}
		result += w
	}

	return result + "\n"
}
