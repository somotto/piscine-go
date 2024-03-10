package piscine

func ConcatParams(args []string) string {
	var result string
	for _, arg := range args {
		result += arg
	}
	return result
}
