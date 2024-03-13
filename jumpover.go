package piscine

func JumpOver(str string) string {
	if str == "" || len(str) < 3 {
		return "\n"
	}
	var result string
	for i, char := range str {
		if (i+1)%3 == 0 {
			result += string(char)
		}
	}
	return result + "\n"
}
