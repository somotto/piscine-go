package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
    args := os.Args[1:]
    if len(args) != 2 {
        z01.PrintRune('\n')
        return
    }

    seen := make(map[rune]bool)
    
    for _, str := range args {
        for _, char := range str {
            if !seen[char] {
                z01.PrintRune(char)
                seen[char] = true
            }
        }
    }
    
    z01.PrintRune('\n')
}