package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	}
	for i := nb; i <= nb+5; i++ {
		nb++
		if IsPrime(nb) {
			return nb
		}
	}
	return 0
}
