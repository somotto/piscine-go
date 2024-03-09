package piscine

func IsPrintable(s string) bool {
	for _, b := range s {
		if b < 32 || b > 126 {
			return false
		}
	}
	return true
}
