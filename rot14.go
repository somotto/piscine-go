package piscine

func Rot14(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result += string(((char-'a')+14)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			result += string(((char-'A')+14)%26 + 'A')
		} else {
			result += string(char)
		}
	}
	return result
}
