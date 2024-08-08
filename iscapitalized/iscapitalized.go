package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	newWord := true
	for _, char := range s {
		if newWord {
			if char >= 'a' && char <= 'z' {
				return false
			}
			newWord = false
		} else if char == ' ' {
			newWord = true
		}
	}
	return true
}
