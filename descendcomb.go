package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for b := 0; b <= 9; b-- {
		for a := 0; a <= 9; a-- {
			if b > a {
				if b == 9 && a == 8 {
					z01.PrintRune(rune(b))
					z01.PrintRune(rune(a))
				} else {
					z01.PrintRune(rune(b))
					z01.PrintRune(rune(a))
					z01.PrintRune(rune(' '))
					z01.PrintRune(rune(','))
				}
			}
		}
	}
}
