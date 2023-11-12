package twenty

import (
	"math"
	"sort"
	"strconv"
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

func TimeConvert(t string) string {
	var result string
	if t[len(t)-2:] == "AM" && t[:2] == "12" {
		result = "00" + t[2:len(t)-2]
	} else if t[len(t)-2:] == "AM" {
		result = t[:len(t)-2]
	} else if t[len(t)-2:] == "PM" && t[:2] == "12" {
		result = t[:len(t)-2]
	} else {
		hour, err := strconv.Atoi(t[:2])
		if err != nil {
			panic("Unable to convert to int.")
		}
		result = strconv.Itoa(12+hour) + t[2:len(t)-2]
	}
	return result
}

func SparseArray(values, queries []string) []int {
	var result []int
	freqs := make(map[string]int)
	for _, x := range values {
		freqs[x] += 1
	}

	for _, q := range queries {
		if val, ok := freqs[q]; ok {
			result = append(result, val)
		} else {
			result = append(result, 0)
		}
	}
	return result
}

func LonelyInt(nums []int) int {
	var result int
	freqs := make(map[int]int)

	for idx := range nums {
		freqs[nums[idx]] += 1
	}

	for key, val := range freqs {
		if val == 1 {
			result = key
			break
		}
	}
	return result
}

func DiagDiff(mat [][]int) float64 {
	j := len(mat) - 1
	td := 0.0
	du := 0.0
	for i := range mat {
		td += float64(mat[i][i])
		du += float64(mat[j][i])
		j -= 1
	}
	return math.Abs(td - du)
}
