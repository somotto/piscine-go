package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func main() {
	if len(os.Args) != 2 {
		// Print nothing if there are more than one or no arguments
		return
	}

	arg := os.Args[1]
	num := 0
	sign := 1
	for i, char := range arg {
		if char == '-' && i == 0 {
			sign = -1
			continue
		}
		if char < '0' || char > '9' {
			// Handle invalid input (non-integer argument)
			printString("Invalid input. Please provide a positive integer.\n")
			return
		}
		num = num*10 + int(char-'0')
	}

	num *= sign

	if isPowerOfTwo(num) {
		printString("true\n")
	} else {
		printString("false\n")
	}
}

func printString(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}

// 	num := atoi(os.Args[1])
// 	if isPowerOf2(num) {
// 		printString("true")
// 	} else {
// 		printString("false")
// 	}
// }

// func atoi(s string) int {
// 	n := 0
// 	for _, ch := range s {
// 		n = n*10 + int(ch-'0')
// 	}
// 	return n
// }

// func isPowerOf2(n int) bool {
// 	return n > 0 && (n&(n-1) == 0)
// }

// func printString(s string) {
// 	for _, ch := range s {
// 		z01.PrintRune(ch)
// 	}
// 	z01.PrintRune('\n')
// }
