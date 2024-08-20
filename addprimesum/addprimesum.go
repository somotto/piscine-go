package main

import (
	"os"

	"github.com/01-edu/z01"
)

func atoi(s string) int {
	num := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}
		num = num*10 + int(c-'0')
	}
	return num
}

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

func itoa(n int) string {
	if n <= 0 {
		return "0"
	}
	var digits []byte
	for n > 0 {
		digits = append(digits, byte(n%10)+'0')
		n /= 10
	}
	for i,j := 0, len(digits) -1; i<j; i,j = i+1,j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return string(digits)
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	n:= atoi(os.Args[1])
	sum := sumPrimes(n)
	printint := itoa(sum)

	for _, v := range printint {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

