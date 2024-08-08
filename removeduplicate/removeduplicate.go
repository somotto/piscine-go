package main

import (
	"fmt"
)

func main() {
	fmt.Println(RemoveDuplicate([]int{1, 2, 3, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(RemoveDuplicate([]int{1, 1, 2, 2, 3}))
	fmt.Println(RemoveDuplicate([]int{}))
}

func RemoveDuplicate(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}

	// Create a map to track seen integers
	seen := make(map[int]bool)
	// Create a slice to store unique integers
	result := []int{}

	for _, value := range slice {
		if !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}

	return result
}
