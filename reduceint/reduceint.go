package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	if len(a) == 0 {
		return
	}
	result := a[0]
	for i := 1; i < len(a); i++ {
		result = f(result, a[i])
	}
	str := strconv.Itoa(result)
	for _, ch := range str {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}