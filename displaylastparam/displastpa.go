package main

import (
	"os"

	"github.com/01-edu/z01"
)


func main() {
	if len(os.Args[1:]) == 0 {
		return
	}
	last := os.Args[len(os.Args)-1]
	for _, ch := range last {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}