package main

import (
	"fmt"
)

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}

func RetainFirstHalf(str string) string {
    length := len(str)
    if length == 0 {
        return ""
    }
    if length == 1 {
        return str
    }
    mid := length / 2
    return str[:mid]
}

