package main

import (
	"github.com/01-edu/z01"
)

func main() {
	// Loop through all possible digits for the first position
	for i := '9'; i >= '0'; i-- {
		// Loop through all possible digits for the second position
		for j := i - 1; j >= '0'; j-- {
			// Loop through all possible digits for the third position
			for k := j - 1; k >= '0'; k-- {
				// Print the combination
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				// If this is not the last combination, print a comma and a space
				if i != '2' || j != '1' || k != '0' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	// Print a newline at the end
	z01.PrintRune('\n')
}
