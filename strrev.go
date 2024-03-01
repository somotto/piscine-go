package piscine

func StrRev(s string) string {
	var reveservedString string
	for i := len(s) - 1; i >= 0; i-- {
		reveservedString += string(s[i])
	}
	return reveservedString
}
