package piscine

func StringToIntSlice(str string) []int {
	res := []int{}
	var num int
	for _, char := range str {
		if '0' <= char && char <= '9' {
			num = num*10 + int(char-'0')
		} else if num > 0 {
			res = append(res, num)
			num = 0
		}
	}
	if num > 0 {
		res = append(res, num)
	}
	return res
}
