package ten

// Use two pointers. One slow, one fast. Perform swaps to the slow index.
func RemoveDups(nums []int) int {
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
