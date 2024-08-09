package main

import (
	"os"
	"github.com/01-edu/z01"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func sumPrimes(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}

	var digits []rune
	for n > 0 {
		digits = append(digits, rune('0'+n%10))
		n /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
	z01.PrintRune('\n')
}

func atoi(s string) (int, bool) {
	num := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, false
		}
		num = num*10 + int(c-'0')
	}
	return num, true
}

func main() {
	if len(os.Args) != 2 {
		printInt(0)
		return
	}

	n, ok := atoi(os.Args[1])
	if !ok || n <= 0 {
		printInt(0)
		return
	}

	printInt(sumPrimes(n))
}

