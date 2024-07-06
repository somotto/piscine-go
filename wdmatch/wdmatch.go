package main

import (
	"os"

	"github.com/01-edu/z01"
)


func main() {
	args := os.Args
	if len(args) != 3 {
		return
	}

	str1, str2 := args[1], args[2]
	i, j := 0, 0
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
		}
		j++
	}

	if i == len(str1) {
		for _, r := range str1 {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
