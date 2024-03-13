package piscine

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	mid := len(nums) / 2
	if len(nums)%2 == 0 {
		return (nums[mid-1] + nums[mid]) / 2
	}
	return nums[mid]
}
