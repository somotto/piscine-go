package main

import (
	"github.com/01-edu/z01"
)

var s string = "x = 42, y = 21"

func main() {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
