package main

import (
	"fmt"
)

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}

func FifthAndSkip(str string) string {
	// Edge cases
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Input\n"
	}

	result := ""
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			result += string(str[i])
			count++
			if count == 5 {
				result += " "
				count = 0
			}
		}
	}

	// Trim trailing space if it exists
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
