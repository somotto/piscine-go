package piscine

func UltimateDivMod(a *int, b *int) {
	// b must be more than 0
	if *b == 0 {
		panic("b cannot be zero")
	}
	div := *a / *b
	mod := *a % *b

	*a = div
	*b = mod
}
