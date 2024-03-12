package main

import (
	"os"

	"github.com/01-edu/z01"
)

const EvenMsg string = "I have an even number of arguments"

const OddMsg string = "I have an odd number of arguments"

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func isEven(nbr int) bool {
	if even(nbr) {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg := len(os.Args[1:])
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
