package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(dec string) string {
	result := ""
	size := len(dec)

	for _, char := range dec {
		hashedChar := (int(char) + size) % 127
		if hashedChar < 33 {
			hashedChar += 33
		}
		result += string(rune(hashedChar))
	}

	return result
}
