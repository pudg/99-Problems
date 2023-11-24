package twenty

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
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

func SortCount(nums []int) {
	maxNum := slices.Max(nums)
	zeros := make([]int, maxNum+1)
	for i := range nums {
		zeros[nums[i]] += 1
	}

	for i := range zeros {
		if zeros[i] != 0 {
			for j := 0; j < zeros[i]; j++ {
				fmt.Println(i)
			}
		}
	}
}

func Pangram(s string) bool {
	s = strings.ToLower(s)
	chars := make(map[string]int)
	for i := range s {
		if s[i] != ' ' {
			chars[string(s[i])] += 1
		}
	}

	return len(chars) == 26
}

func ArrayPerm(A, B []int, k int) string {
	sort.Ints(A)
	sort.Slice(B, func(i, j int) bool {
		return B[i] > B[j]
	})

	for i := range A {
		if A[i]+B[i] < k {
			return "NO"
		}
	}

	return "YES"
}

func SubarrayDivision(nums []int, day, month int) int {
	result := 0

	for i := range nums {
		subArr := nums[i : i+month]
		if SumIntSlice(subArr) == day {
			result += 1
		}
	}
	return result
}

func SalesMatch(nums []int) int {
	pairs := 0
	freqs := make(map[int]int)

	for i := range nums {
		freqs[nums[i]] += 1
		if freqs[nums[i]]%2 == 0 {
			pairs += 1
		}
	}
	return pairs
}

func BookPage(num, page int) int {
	lhs := page / 2
	rhs := (num / 2) - (page / 2)
	return min(lhs, rhs)
}

func UnfairMin(arr []int, k int) int {
	sort.Ints(arr)
	globalMin := arr[len(arr)-1]

	for i := 0; i < len(arr)-k+1; i++ {
		diff := arr[i+k-1] - arr[i]
		globalMin = min(globalMin, diff)
	}
	return globalMin
}

func DynamicArray(n int, queries [][]int) []int {
	arr := make([][]int, n)
	answers := []int{}
	lastAnswer := 0

	for i := range queries {
		if queries[i][0] == 1 {
			idx := ((queries[i][1] ^ lastAnswer) % n)
			arr[idx] = append(arr[idx], queries[i][2])
		} else {
			idx := ((queries[i][1] ^ lastAnswer) % n)
			lastAnswer = arr[idx][queries[i][2]%len(arr[idx])]
			answers = append(answers, lastAnswer)
		}
	}
	return answers
}

func NearestPower2(n int) int {
	return int(math.Pow(2, math.Floor(math.Log2(float64(n)))))
}

func CounterGame(n int) string {
	move := true

	for {
		nearest := NearestPower2(n)
		if n != nearest {
			n = n - nearest
		} else {
			n = n / 2
		}
		if n == 1 {
			break
		}
		move = !move
	}
	if move {
		return "Player1"
	} else {
		return "Player2"
	}
}

type Node struct {
	val  string
	next *Node
}

func ReverseLL(root *Node) {
	curr := root
	var prev *Node
	prev = nil

	for {
		currNext := curr.next
		curr.next = prev
		prev = currNext
		curr = currNext
		if curr == nil {
			break
		}
	}
}
