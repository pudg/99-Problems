package twenty

import (
	"sort"
)

func MapValues[K comparable, V any](values map[K]V) []V {
	vals := make([]V, 0, len(values))

	for _, val := range values {
		vals = append(vals, val)
	}
	return vals
}

func PlusMinus(nums []int) []float32 {
	var freqs []float32

	counts := make(map[int]int)

	for i := range nums {
		if nums[i] > 0 {
			counts[1] += 1
		} else if nums[i] < 0 {
			counts[-1] += 1
		} else {
			counts[0] += 1
		}
	}
	values := MapValues(counts)
	for _, val := range values {
		freqs = append(freqs, float32(val)/float32(len(nums)))
	}
	return freqs
}

func SumIntSlice(nums []int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func MinMaxSum(nums []int) []int {
	var sums []int
	sort.Ints(nums)
	min := SumIntSlice(nums[0:4])
	max := SumIntSlice(nums[len(nums)-4:])
	sums = append(sums, min, max)
	return sums
}
