package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	if nb%2 == 0 {
		nb++
	}
	for !IsPrime(nb) {
		nb += 2
	}
	return nb
}
