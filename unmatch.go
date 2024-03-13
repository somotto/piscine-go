package piscine

func Unmatch(a []int) int {
	counts := make([]int, len(a))
	for _, num := range a {
		counts[num]++
	}
	for i, count := range counts {
		if count%2 != 0 {
			return a[i]
		}
	}
	return -1
}
