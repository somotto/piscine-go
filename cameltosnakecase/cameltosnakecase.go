package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := ""
	valid := true
	for i := 0; i < len(s); i++ {
		if (s[i] < 'A' || s[i] > 'Z') && (s[i] < 'a' || s[i] > 'z') {
			valid = false
			break
		}
		if i > 0 && s[i] >= 'A' && s[i] <= 'Z' {
			if (i < len(s)-1 && s[i+1] >= 'a' && s[i+1] <= 'z') && 
			   (s[i-1] >= 'a' && s[i-1] <= 'z') {
				result += "_"
			} else if s[i-1] >= 'A' && s[i-1] <= 'Z' {
				valid = false
				break
			}
		}
		result += string(s[i])
	}

	if !valid || result == s {
		return s
	}

	return result
}