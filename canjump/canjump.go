package main

import (
	"fmt"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}

func CanJump(nums []uint) bool {
	if len(nums) == 0 {
		return false
	}

	lastPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+int(nums[i]) >= lastPos {
			lastPos = i
		}
	}

	return lastPos == 0
}
