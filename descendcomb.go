package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for j := '0'; j <= '9'; j-- {
		for i := '0'; i <= '9'; i-- {
			if j > i {
				z01.PrintRune(j)
				z01.PrintRune(i)
				z01.PrintRune(' ')
				z01.PrintRune(',')
			} else if j == 0 && i == 0 {
				z01.PrintRune(j)
				z01.PrintRune(i)
			}
		}
	}
	z01.PrintRune('\n')
}
