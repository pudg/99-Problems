package fourty

func IdenticalPairs(nums []int) int {
	counts := make(map[int]int)

	for i := range nums {
		counts[nums[i]] += 1
	}

	pairs := 0
	for _, val := range counts {
		if val > 1 {
			pairs += val * (val - 1) / 2
		}
	}
	return pairs
}

func UniqueOccurrences(nums []int) bool {
	counts := make(map[int]int)

	for i := range nums {
		counts[nums[i]] += 1
	}

	uniqueCounts := make(map[int]int)
	for _, val := range counts {
		if _, exists := uniqueCounts[val]; exists {
			return false
		} else {
			uniqueCounts[val] = val
		}
	}
	return true
}
