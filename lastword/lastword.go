package main

import (
	"github.com/01-edu/z01"
)

func LastWord(s string) {
    lastWordStart := -1
    wordEnd := -1

    // Find the last non-space character
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] != ' ' {
            wordEnd = i
            break
        }
    }

    // If no non-space character found, return empty line
    if wordEnd == -1 {
        z01.PrintRune('\n')
        return
    }

    // Find the start of the last word
    for i := wordEnd; i >= 0; i-- {
        if s[i] == ' ' {
            lastWordStart = i + 1
            break
        }
    }

    // If no space found, the whole string is the last word
    if lastWordStart == -1 {
        lastWordStart = 0
    }

    // Print the last word
    for i := lastWordStart; i <= wordEnd; i++ {
        z01.PrintRune(rune(s[i]))
    }
    z01.PrintRune('\n')
}

func main() {
    LastWord("this        ...       is sparta, then again, maybe    not")
    LastWord(" lorem,ipsum ")
    LastWord(" ")
}