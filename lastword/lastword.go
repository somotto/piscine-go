package main

import (
	"fmt"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

func LastWord(s string) string {
	lastWord := ""
	inWord := false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			inWord = true
			lastWord = string(s[i]) + lastWord
		} else if inWord {
			break
		}
	}

	return lastWord + "\n"
}
