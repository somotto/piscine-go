package main

import "fmt"

// WordFlip reverses the order of words in a string
func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output" + "\n"
	}
	result := ""
	slice := []string{}
	for _, ch := range str {
		if ch != ' ' {
			result += string(ch)
		} else if result != "" {
			slice = append(slice, result)
			result = ""
		}
	}
	if result != "" {
		slice = append(slice, result)
	}
	wordflipped := ""
	for i := len(slice)-1; i>= 0; i-- {
		if i != 0 {
			wordflipped += slice[i] + " "
		} else {
			wordflipped += slice[i]
		}
	}
	return wordflipped + "\n"
}

func main() {
	// Example test cases
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
