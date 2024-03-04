package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	for i := 1; i <= nb; i++ {
		if result > (1 << 31) {
			return 0
		}
		result *= i
	}
	return result
}
