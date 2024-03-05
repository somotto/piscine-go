package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	if nb%2 == 0 {
		nb++
	}
	for i := nb + 2; i < nb/2; i += 2 {
		if IsPrime(i) {
			return i
		}
	}
	if IsPrime(nb) {
		return nb
	}
	return 0
}
