package piscine

func ToUpper(s string) string {
	result := ""
	for _, b := range s {
		if 'a' <= b && b <= 'z' {
			result += string(b - 32)
		} else {
			result += string(b)
		}
	}
	return result
}
