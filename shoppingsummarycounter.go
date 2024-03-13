package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""
	for _, char := range str {
		if char == ' ' || char == '\t' || char == '\n' {
			if word != "" {
				summary[word] = summary[word] + 1
			}
			word = ""
		} else {
			word += string(char)
		}
	}
	if word != "" {
		summary[word] = summary[word] + 1
	}
	summary[":3"] = summary[":3"] + 1
	return summary
}
