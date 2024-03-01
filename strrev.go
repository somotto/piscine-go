package piscine

func StrRev(s string) string {
	r := []rune(s)
	length := len(s)
	var revRun []rune

	for i := length - 1; i >= 0; i-- {
		revRun = appen(revRun, r[i])
	}
	return string(revRun)
}
