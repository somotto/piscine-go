package piscine

func Capitalize(s string) string {
	word := []rune(s)
	for i, char := range word {
		if isNumber(char) {
			if i == 0 || isNumber(word[i-1]) == false {
				if word[i] >= 'a' && word[i] <= 'z' {
					word[i] = char - 32
				}
			} else {
				if word[i] >= 'A' && word[i] <= 'Z' {
					word[i] = char + 32
				}
			}
		}
	}
	return string(word)
}

func isNumber(r rune) bool {
	if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
		return true
	}
	return false
}
