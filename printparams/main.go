package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args[1:]
	for i := 0; i < len(params); i++ {
		for _, char := range params[i] {
			z01.PrintRune(rune(char))
		}
		z01.PrintRune('\n')
	}
}
