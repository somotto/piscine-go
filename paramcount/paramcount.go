package main

import (
	"os"

	"github.com/01-edu/z01"
)


func main() {
argcount := len(os.Args)-1
if argcount == 0 {
	z01.PrintRune('0')
} else {
	printNumber(argcount)
}
z01.PrintRune('\n')
}
func printNumber(n int) {
	if n >= 10 {
		// Recursive call
		printNumber(n/10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}