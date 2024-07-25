package main

import (
	"fmt"
)

func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}

func CheckNumber(arg string) bool {
	for _, v := range arg {
		if v >= '0' && v <= '9' {
			return true
		}
	}
	return false
}
