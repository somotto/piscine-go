package main

import (
	"os"
)

func main() {
	// Validate the number of arguments
	if len(os.Args) != 4 {
		return
	}

	// Parse input values (assuming they are valid integers)
	value1 := atoi(os.Args[1])
	value2 := atoi(os.Args[3])

	if value1 >= 9223372036854775807 {
		return
	}
	if value1 <= -9223372036854775807 {
		return
	}

	// Perform the operation based on the operator
	operator := os.Args[2]
	var result int

	if operator == "+" {
		result = value1 + value2
	} else if operator == "-" {
		result = value1 - value2
	} else if operator == "*" {
		result = value1 * value2
	} else if operator == "/" {
		if value2 != 0 {
			result = value1 / value2
		} else {
			os.Stderr.WriteString("Error: Division by zero\n")
			os.Exit(1)
		}
	} else if operator == "%" {
		if value2 != 0 {
			result = value1 % value2
		} else {
			os.Stderr.WriteString("Error: Modulo by zero\n")
			os.Exit(1)
		}
	} else {
		os.Stderr.WriteString("Error: Invalid operator\n")
		os.Exit(1)
	}

	// Convert the result back to a string (assuming it fits within int range)
	resultStr := itoa(result)

	// Print the result
	os.Stdout.Write([]byte(resultStr + "\n"))
}

// atoi converts a string to an integer (simplified implementation)
func atoi(s string) int {
	n := 0
	sign := 1

	if s[0] == '-' {
		sign = -1
		s = s[1:]
	}

	for _, c := range []byte(s) {
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		} else {
			break
		}
	}

	return sign * n
}

// itoa converts an integer to a string (simplified implementation)
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte(n%10) + '0'
		n /= 10
	}
	return string(buf[i:])
}
