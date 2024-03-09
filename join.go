package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	concatstring := strs[0]
	for _, ele := range strs[1:] {
		concatstring += sep + ele
	}
	return concatstring
}
