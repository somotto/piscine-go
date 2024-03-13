package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	LengthOfArgs := len(os.Args[1:])
	if LengthOfArgs == 0 {
		fmt.Println("File name missing")
		return
	} else if LengthOfArgs > 1 {
		fmt.Println("Too many arguments")
		return
	}
	filename := os.Args[1]
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Printf("The Filename Is missing%v\n", filename)
		return
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf(string(content))
}
