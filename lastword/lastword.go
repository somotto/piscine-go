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
    // Trim trailing spaces
    s = trimTrailingSpaces(s)

    // If the string is empty after trimming, return a newline
    if s == "" {
        return "\n"
    }

    // Find the last word
    lastSpaceIndex := -1
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == ' ' {
            lastSpaceIndex = i
            break
        }
    }

    // Return the last word followed by a newline
    return s[lastSpaceIndex+1:] + "\n"
}

func trimTrailingSpaces(s string) string {
    end := len(s) - 1
    for end >= 0 && s[end] == ' ' {
        end--
    }
    return s[:end+1]
}