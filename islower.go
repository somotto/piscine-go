package piscine

func IsLower(s string) bool {
	for _, r := range s {
		if r >= 'A'-'Z' {
			return false
		}
	}
	return true
}
