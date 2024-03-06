package piscine

func LastRune(s string) rune {
	if s == "" {
		return 0
	}
	return []rune(s)[len(s)-1]
}
