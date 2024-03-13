package piscine

func Unmatch(a []int) int {
	for _, char := range a {
		b := 0
		for _, d := range a {
			if d == char {
				b++
			}
		}
		if b == 1 || b%2 == 1 {
			return char
		}
	}
	return -1
}
