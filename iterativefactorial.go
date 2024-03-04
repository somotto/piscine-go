package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb < 0 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		if result > (1 << 31) {
			return 0
		}
		result *= i
	}
	return result
}
