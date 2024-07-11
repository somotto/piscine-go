package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	str := os.Args[1]
	oldChar := os.Args[2]
	newChar := os.Args[3]

	for _, char := range str {
		if string(char) == oldChar {
			for _, ch := range str {
				if string(ch) == oldChar {
					z01.PrintRune(rune(newChar[0]))
				} else {
					z01.PrintRune(ch)
				}
			}
			z01.PrintRune('\n')
			return
		}
	}

	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
