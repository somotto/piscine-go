package piscine

func ConcatParams(args []string) string {
	var result string
	for i, value := range args {
		if i > 0 {
			result += value + "\n"
		}
	}
	return result
}
