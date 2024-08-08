package main

import (
	"fmt"
)

func main() {
	fmt.Print(ReverseSecondHalf("This is the 1st half This is the 2nd half"))
	fmt.Print(ReverseSecondHalf(""))
	fmt.Print(ReverseSecondHalf("Hello World"))
}

func ReverseSecondHalf(str string) string {
    length := len(str)
    
    if length == 0 {
        return "Invalid Output\n"
    }
    
    if length == 1 {
        return str + "\n"
    }

    // Calculate the start index for the second half
    start := (length + 1) / 2 // Round up for odd length

    // Get the second half of the string
    secondHalf := str[start:]

    // Reverse the second half
    reversed := ""
    for i := len(secondHalf) - 1; i >= 0; i-- {
        reversed += string(secondHalf[i])
    }

    return reversed + "\n"
}
