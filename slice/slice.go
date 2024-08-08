package main

import (
	"fmt"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return nil
	}

	start := 0
	end := len(a)

	// Handle the start index
	if len(nbrs) >= 1 {
		start = nbrs[0]
		if start < 0 {
			start = len(a) + start
		}
		if start < 0 {
			start = 0
		}
		if start > len(a) {
			start = len(a)
		}
	}

	// Handle the end index
	if len(nbrs) >= 2 {
		end = nbrs[1]
		if end < 0 {
			end = len(a) + end
		}
		if end < 0 {
			end = 0
		}
		if end > len(a) {
			end = len(a)
		}
		if end < start {
			return nil
		}
	}

	// Return the slice based on the calculated start and end indices
	return a[start:end]
}
