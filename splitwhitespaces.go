package piscine

func SplitWhiteSpaces(s string) []string {
	var value []string
	var word string
	for _, char := range s {
		if char == '\n' || char == '\t' || char == ' ' {
			if word != "" {
				value = append(value, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		value = append(value, word)
	}
	return value
}
