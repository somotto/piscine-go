// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(Itoa(12345))
// 	fmt.Println(Itoa(0))
// 	fmt.Println(Itoa(-1234))
// 	fmt.Println(Itoa(987654321))
// }

// func Itoa(n int) string {
// 	// handle zero
// 	if n == 0 {
// 		return "0"
// 	}

// 	// handle negative numbers
// 	negative := false
// 	if n < 0 {
// 		negative = true
// 		n = -n
// 	}

// 	// convert digits to charactrs
// 	var digits []byte
// 	for n > 0 {
// 		digits = append(digits, byte(n%10)+'0')
// 		n /= 10
// 	}

// 	// Reverse the digits
// 	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
// 		digits[i], digits[j] = digits[j], digits[i]
// 	}

// 	if negative {
// 		digits = append([]byte{'-'}, digits...)
// 	}

// 	return string(digits)
// }

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {
	result := 0
	sign := 1
	start := 0

	if len(s) == 0 {
		return 0
	}

	// Handle negative numbers
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		digit := int(s[i] - '0')
		result = result*10 + digit
	}

	return sign * result
}
