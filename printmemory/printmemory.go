package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	for i, v := range arr {
		hex := hextodec(int(v))
		if len(hex) == 1 {
			z01.PrintRune('0')
		}
		for _, ch := range hex {
			z01.PrintRune(rune(ch))
		}
		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		} else if i != 9 {
			z01.PrintRune(' ')
		}
	}
	for _, v := range arr {
		if v >= 32 && v <= 126 {
			z01.PrintRune(rune(v))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func hextodec(dec int) string {
	if dec == 0 {
		return "00"
	}
	s := "0123456789abcdef"
	s1 := ""
	for dec > 0 {
		s1 = string(s[dec%16]) + s1
		dec /= 16
	}
	return s1
}