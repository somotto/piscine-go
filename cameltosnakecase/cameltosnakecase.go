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
	for i := 0; i < len(s); i++ {
		if (i == len(s)-1 && (s[i] >= 'A' && s[i] <= 'Z')) || s[i] >= 'A' && s[i] <= 'Z' && s[i+1] >= 'A' && s[i+1] <= 'Z' {
			return s
		}
		if i < len(s) && !((s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z')) {
			return s
		}

		if i > 0 && s[i] >= 'A' && s[i] <= 'Z' && (s[i-1] >= 'a' && s[i-1] <= 'z') {
			result += "_"
		}
		result += string(s[i])
	}
	return result
}
