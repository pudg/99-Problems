package ten

import (
	"fmt"
	"math"
	"regexp"
	"sort"
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

func SortString(s string) string {
	result := strings.Split(s, "")
	sort.Strings(result)
	return strings.Join(result, "")
}

func GroupAnagrams(words []string) [][]string {
	anagrams := make(map[string][]string)

	for _, word := range words {
		sortedWord := SortString(word)
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}

	result := make([][]string, 0, len(anagrams))
	for _, val := range anagrams {
		result = append(result, val)
	}
	fmt.Println(result)
	return result
}

func TopKFreq(nums []int, k int) []int {
	freqs := make(map[int]int)

	for _, num := range nums {
		freqs[num]++
	}

	revFreqs := make(map[int][]int)

	for key, val := range freqs {
		revFreqs[val] = append(revFreqs[val], key)
	}

	result := []int{}

	for i := len(nums); len(result) != k; i-- {
		for _, val := range revFreqs[i] {
			if len(result) != k {
				result = append(result, val)
			}
		}
	}

	return result
}
