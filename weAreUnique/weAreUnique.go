package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	// Create maps to store unique characters
	chars1 := make(map[rune]bool)
	chars2 := make(map[rune]bool)

	// Add characters from str1 to chars1
	for _, char := range str1 {
		chars1[char] = true
	}

	// Add characters from str2 to chars2
	for _, char := range str2 {
		chars2[char] = true
	}

	uniqueCount := 0

	// Count characters unique to str1
	for char := range chars1 {
		if !chars2[char] {
			uniqueCount++
		}
	}

	// Count characters unique to str2
	for char := range chars2 {
		if !chars1[char] {
			uniqueCount++
		}
	}

	return uniqueCount
}
