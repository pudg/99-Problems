package ten

import (
	"math"
	"regexp"
	"strings"
)

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

func TwoSum(nums []int, target int) []int {
	locations := make(map[int][]int)

	for idx, num := range nums {
		locations[num] = append(locations[num], idx)
	}

	for idx, num := range nums {
		delta := target - num
		if _, exists := locations[delta]; exists {
			deltaIdx := locations[delta][0]
			if num+nums[deltaIdx] == target {
				if idx == deltaIdx {
					if len(locations[delta]) >= 2 {
						return locations[delta][:2]
					}
				} else {
					return []int{idx, deltaIdx}
				}
			}
		}
	}

	return []int{}
}

// Use a hash map to store index locations of values.
// Finds solution in a single pass through the slice of nums.
func TwoSumV2(nums []int, target int) []int {
	locations := make(map[int]int)

	for numIdx, num := range nums {
		delta := target - num
		if deltaIdx, exists := locations[delta]; exists {
			return []int{numIdx, deltaIdx}
		} else {
			locations[num] = numIdx
		}
	}

	return []int{}
}

// Use two pointers and work your way inwards.
func ValidPalindrome(phrase string) bool {
	re := regexp.MustCompile("[^a-zA-Z0-9]")
	answer := re.ReplaceAllString(phrase, "")
	answer = strings.ToLower((answer))
	lhs := 0
	rhs := len(answer) - 1
	for lhs != rhs {
		if string(answer[lhs]) != string(answer[rhs]) {
			return false
		}
		lhs += 1
		rhs -= 1
	}
	return true
}

// Use Newtons Method to approximate square root.
func NewtonSquareRoot(num int32, tol float64) float64 {
	x := float64(num)
	root := float64(0.5 * (x + (float64(num) / x)))
	for !(math.Abs(float64(x-root)) < tol) {
		x = root
		root = float64(0.5 * (x + (float64(num) / x)))
	}
	return root
}

// Use a hashmap to memorize previously seen integers.
func ContainsDup(nums []int) bool {
	memory := make(map[int]int)

	for _, num := range nums {
		if _, exists := memory[num]; exists {
			return true
		}
		memory[num] = 1
	}

	return false
}

func ValidAnagram(s string, t string) bool {
	return false
}
