package main

import (
	"os"

	"github.com/01-edu/z01"
)

var romanNumerals = []struct {
	value  int
	symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func main() {
	if len(os.Args) != 2 {
		printError()
		return
	}

	num := atoi(os.Args[1])
	if num <= 0 || num >= 4000 {
		printError()
		return
	}

	calculation := ""
	roman := ""

	for _, rn := range romanNumerals {
		for num >= rn.value {
			if len(calculation) > 0 {
				calculation += "+"
			}
			if len(rn.symbol) == 2 {
				calculation += "(" + string(rn.symbol[1]) + "-" + string(rn.symbol[0]) + ")"
			} else {
				calculation += rn.symbol
			}
			roman += rn.symbol
			num -= rn.value
		}
	}

	printString(calculation)
	z01.PrintRune('\n')
	printString(roman)
	z01.PrintRune('\n')
}

func atoi(s string) int {
	n := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0
		}
		n = n*10 + int(ch-'0')
	}
	return n
}

func printString(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func printError() {
	errorMsg := "ERROR: cannot convert to roman digit"
	printString(errorMsg)
	z01.PrintRune('\n')
}