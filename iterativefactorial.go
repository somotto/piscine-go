package piscine

func IterativeFactorial(nb int) int {
	if nb > 0 && nb <= 20 {
		result := 1
		for i := 1; i <= nb; i++ {
			result *= i
			return result
		}
	} else if nb == 0 {
		return 1
	}
	return 0
}
