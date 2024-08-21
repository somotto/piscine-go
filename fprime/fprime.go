package main

import (
	"os"

	"github.com/01-edu/z01"
)

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


// prints an integer digit by digit using z01.PrintRune
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	negative := false 
	if n < 0 {
		negative = true
		n = -n
	}
	var digits []byte
	for n > 0 {
		digits = append(digits, byte(n%10) + '0')
		n/=10
	}
	for i,j := 0,len(digits)-1;i < j;i,j = i+1,j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	} 
	if negative {
		digits = append([]byte{'-'}, digits...)
	}
	return string(digits)
}
//read command line args
//convert commandline arg(string) to int(Atoi)
// Isprime func then Primefactors
// if factors are more than 0, range over them printing * in between
// then convert fprimes(int) to str(Itoa), print them
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
			number:= itoa(factor)
			for _,v := range number {
				z01.PrintRune(v)
			}
		}
		z01.PrintRune('\n')
	}
}
