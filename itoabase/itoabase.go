package main

import (
	"fmt"
)

func main() {
	fmt.Println(ItoaBase(10, 2))
	fmt.Println(ItoaBase(255, 16))
	fmt.Println(ItoaBase(-42, 4))
	fmt.Println(ItoaBase(123, 10))
	fmt.Println(ItoaBase(0, 8))
	fmt.Println(ItoaBase(255, 2))
	fmt.Println(ItoaBase(-255, 16))
	fmt.Println(ItoaBase(15, 16))
	fmt.Println(ItoaBase(10, 4))
	fmt.Println(ItoaBase(255, 10))
}

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	if value == 0 {
		return "0"
	}

	digits := "0123456789ABCDEF"
	result := ""
	isNegative := value < 0

	if isNegative {
		value = -value
	}

	for value > 0 {
		remainder := value % base
		result = string(digits[remainder]) + result
		value = value / base
	}

	if isNegative {
		result = "-" + result
	}

	return result
}
