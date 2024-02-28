package main

import "github.com/01-edu/z01"

func main() {
	var x rune = 122
	for i := x; i <= 97; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
