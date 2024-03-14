package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output"
	}
	result := ""
	i := 0
	for i < len(str) {
		word := ""
		for j := 0; j < 5 && i+j < len(str); j++ {
			word += string(str[i+j])
		}
		if len(word) == 5 {
			result += word + "\n"
			i += 6
		} else {
			result += word + "\n"
			i = len(str)
		}
	}
	return result
}
