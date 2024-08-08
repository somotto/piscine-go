package main

import "fmt"

func main() {
	fmt.Println(ReverseStrings([]string{"a", "b", "c"}))
	fmt.Println(ReverseStrings([]string{"Good", "Morning!"}))
	fmt.Println(ReverseStrings([]string{"Hello World"}))
}

func ReverseStrings(strs []string) string {
    result := ""
    
    // Iterate over the slice in reverse order
    for i := len(strs) - 1; i >= 0; i-- {
        // Reverse the current string
        reversedStr := reverseString(strs[i])
        // Append the reversed string to the result
        result += reversedStr
        // Append a space after each element except the last one
        if i > 0 {
            result += " "
        }
    }
    
    return result
}

// Helper function to reverse a string
func reverseString(s string) string {
    reversed := ""
    for j := len(s) - 1; j >= 0; j-- {
        reversed += string(s[j])
    }
    return reversed
}

