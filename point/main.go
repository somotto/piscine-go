package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var s string = "x = 42, y = 21"
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
