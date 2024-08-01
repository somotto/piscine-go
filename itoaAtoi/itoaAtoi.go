package main

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	var digits []byte
	for n > 0 {
		digits = append(digits, byte(n%10)+'0')
		n /= 10
	}

	// Reverse the digits
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	if negative {
		digits = append([]byte{'-'}, digits...)
	}

	return string(digits)
}

func Atoi(s string) int {
	result := 0
	sign := 1
	start := 0

	if len(s) == 0 {
		return 0
	}

	// Handle negative numbers
	if s[0] == '-' {
		sign = -1
		start = 1
	}

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		digit := int(s[i] - '0')
		result = result*10 + digit
	}

	return sign * result
}
