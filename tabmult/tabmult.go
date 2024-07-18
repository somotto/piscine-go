package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num := 0
	for _, r := range os.Args[1] {
		if r >= '0' && r <= '9' {
			num = num*10 + int(r-'0')
		} else {
			return
		}
	}

	if num <= 0 || num > 1000000000 {
		return
	}

	for i := 1; i <= 9; i++ {
		result := i * num
		printNumber(i)
		z01.PrintRune(' ')
		z01.PrintRune('x')
		z01.PrintRune(' ')
		printNumber(num)
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		printNumber(result)
		z01.PrintRune('\n')
	}
}

func printNumber(n int) {
	if n < 10 {
		z01.PrintRune(rune(n) + '0')
	} else {
		printNumber(n / 10)
		z01.PrintRune(rune(n%10) + '0')
	}
}
