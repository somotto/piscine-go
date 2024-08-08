package main

import (
	"fmt"
)

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}

// SaveAndMiss processes the input string based on the provided integer
func SaveAndMiss(arg string, num int) string {
	// If num is less than or equal to 0, return the original string
	if num <= 0 {
		return arg
	}

	// If the input string is empty, return an empty string
	if len(arg) == 0 {
		return ""
	}

	result := ""
	length := len(arg)
	index := 0

	// Iterate through the string in steps of 'num'
	for index < length {
		// Save the segment of size 'num'
		end := index + num
		if end > length {
			end = length
		}
		result += arg[index:end]

		// Move the index forward by '2*num' to skip the next segment
		index += 2 * num
	}

	return result
}
