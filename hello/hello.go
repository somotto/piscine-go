package main

import "github.com/01-edu/z01"

func main() {
	str := "Hello World!"
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
