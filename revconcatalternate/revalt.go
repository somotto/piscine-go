package main

import (
	"fmt"
)

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}

func RevConcatAlternate(slice1, slice2 []int) []int {
	newarr := []int{} // Step 1: Initialize an empty slice to store the result

	// Step 2: Handle the case where the slices are of equal length
	if len(slice1) == len(slice2) {
		for i := len(slice1) - 1; i >= 0; i-- { // Loop from the end to the beginning of the slices
			newarr = append(newarr, slice1[i]) // Add the element from slice1
			newarr = append(newarr, slice2[i]) // Add the element from slice2
		}
	}

	// Step 3: Handle the case where slice1 is longer than slice2
	if len(slice1) > len(slice2) {
		for i := len(slice1) - 1; i >= 0; i-- { // Loop from the end to the beginning of slice1
			if i >= len(slice2) {
				newarr = append(newarr, slice1[i]) // Add the remaining elements from slice1 when slice2 is exhausted
			} else {
				newarr = append(newarr, slice1[i]) // Add the element from slice1
				newarr = append(newarr, slice2[i]) // Add the element from slice2
			}
		}
	}

	// Step 4: Handle the case where slice2 is longer than slice1
	if len(slice2) > len(slice1) {
		for i := len(slice2) - 1; i >= 0; i-- { // Loop from the end to the beginning of slice2
			if i >= len(slice1) {
				newarr = append(newarr, slice2[i]) // Add the remaining elements from slice2 when slice1 is exhausted
			} else {
				newarr = append(newarr, slice1[i]) // Add the element from slice1
				newarr = append(newarr, slice2[i]) // Add the element from slice2
			}
		}
	}

	return newarr // Step 5: Return the new slice with the alternated values in reverse order
}
