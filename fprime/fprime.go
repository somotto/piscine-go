package main

import (
	"os"

	"github.com/01-edu/z01"
)

// checks for prime number
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// computes the prime factors of a given number and returns a slice
// if the number itself is prmie, it is returned as the only element
func primeFactors(n int) []int {
	factors := make([]int, 0)
	for i := 2; i <= n; i++ {
		for n%i == 0 && isPrime(i) {
			factors = append(factors, i)
			n /= i
		}
	}
	if len(factors) == 0 {
		factors = append(factors, n)
	}
	return factors
}

// converts string to an interger
func atoi(s string) int {
	res := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return -1
		}
		res = res*10 + int(ch-'0')
	}
	return res
}

// prints an integer digit by digit using z01.PrintRune
func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	digits := []rune{}
	for n > 0 {
		digits = append([]rune{rune(n%10 + '0')}, digits...)
		n /= 10
	}
	for _, digit := range digits {
		z01.PrintRune(digit)
	}
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	num := atoi(os.Args[1])
	if num <= 1 {
		return
	}

	factors := primeFactors(num)

	if len(factors) > 0 {
		for i, factor := range factors {
			if i != 0 {
				z01.PrintRune('*')
			}
			printNumber(factor)
		}
		z01.PrintRune('\n')
	}
}
