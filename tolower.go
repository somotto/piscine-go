package piscine

func ToLower(s string) string {
	result := ""
	for _, b := range s {
		if b >= 'A' && b <= 'Z' {
			result += string(b + 32)
		} else {
			result += string(b)
		}
	}
	return result
}
