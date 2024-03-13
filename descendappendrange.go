package piscine

func DescendAppendRange(max, min int) []int {
	result := []int{}
	if min > max {
		return []int{}
	}
	for i := max; i > min; i-- {
		result = append(result, i)
	}
	return result
}
