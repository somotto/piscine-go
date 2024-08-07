package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func FoldInt(f func(int, int) int, a []int, n int) {
	result := n
	for _, v := range a {
		result = f(result, v)
	}

	//convert the result to a string and print it
	printInt(result)

	//Print a newline at the end
	z01.PrintRune('\n')
}

// Helper function to print an integer
func printInt(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Convert integer to slice of digits
	var digits []int
	for n > 0 {
		digits = append([]int{n % 10}, digits...)
		n /= 10
	}

	// Print each digit
	for _, d := range digits {
		z01.PrintRune(rune('0' + d))
	}
}

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Sub(a, b int) int {
	return a - b
}
