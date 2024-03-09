package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	for _, char := range s {
		if char >= '0' && char <= '9' {
			char = char - '0'
			result = result*10 + int(char)
		} else if result == 0 && char == '-' {
			sign = -1
		}
	}
	return sign * result
}
