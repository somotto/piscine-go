package piscine

func Unmatch(a []int) int {
	for i := range a {
		result := -1
		for j := range a {
			if i != j && a[i] == a[j] {
				result = i
				break
			}
		}
		if result != -1 {
			return a[result]
		}
	}
	return -1
}
