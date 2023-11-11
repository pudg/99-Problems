package twenty

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
