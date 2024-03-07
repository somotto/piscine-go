package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return true
	}
	for _, r := ramge s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return false 
		}
	}
	return true
}
