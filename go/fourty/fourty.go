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

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseLL(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		nxt := curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
	}
	return prev
}

func MergeTwoLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	helper := dummy
	for list1 != nil || list2 != nil {
		if list1 == nil {
			helper.Next = &ListNode{list2.Val, nil}
			helper = helper.Next
			list2 = list2.Next
		} else if list2 == nil {
			helper.Next = &ListNode{list1.Val, nil}
			helper = helper.Next
			list1 = list1.Next
		} else {
			if list1.Val <= list2.Val {
				helper.Next = &ListNode{list1.Val, nil}
				helper = helper.Next
				list1 = list1.Next
			} else {
				helper.Next = &ListNode{list2.Val, nil}
				helper = helper.Next
				list2 = list2.Next
			}
		}
	}
	return dummy.Next
}
