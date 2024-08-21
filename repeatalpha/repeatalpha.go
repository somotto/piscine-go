package main

import (
	"fmt"
)

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}

func RepeatAlpha(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			count := int(char-'a') + 1
			for i := 0; i< count; i++ {
				result += string(char)
			}
		} else 	if char >= 'A' && char <= 'Z' {
			count := int(char-'A') + 1
			for i := 0; i< count; i++ {
				result += string(char)
			}
		} else {
			result += string(char)
		}

	}
	return result
}
