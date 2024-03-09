package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	result := make([]int, 0, max-min)
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}
