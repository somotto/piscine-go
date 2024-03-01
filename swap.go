package piscine

func Swap(a *int, b *int) {
	tempA := *a
	tempB := *b
	*a = tempB
	*b = tempA
}
