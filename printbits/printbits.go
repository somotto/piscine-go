package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num := os.Args[1]
	for _, digit := range num {
		if digit < '0' || digit > '9' {
			printBinary("00000000")
			return
		}
	}

	printBinary(num)
}

func printBinary(num string) {
	value := 0
	for _, digit := range num {
		value = value*10 + int(digit-'0')
	}
	for i := 7; i >= 0; i-- {
		bit := (value >> uint(i)) & 1
		z01.PrintRune(rune(bit) + '0')
	}
}
