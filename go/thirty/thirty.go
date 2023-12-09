package thirty

import "slices"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func BFS(root *Node, target int) bool {
	if root == nil {
		return false
	}
	queue := make([]Node, 0)
	queue = append(queue, *root)
	for len(queue) != 0 {
		curr := queue[0]
		if curr.Left != nil {
			queue = append(queue, *curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, *curr.Right)
		}
		queue = queue[1:]
		if curr.Val == target {
			return true
		}
	}
	return false
}

func MinHeight(root *Node) int {
	if root == nil {
		return 0
	}

	depth := 1
	queue := []*Node{root}

	for len(queue) > 0 {
		queueLen := len(queue)

		for i := 0; i < queueLen; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr.Left == nil && curr.Right == nil {
				return depth
			}
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		depth += 1
	}
	return depth
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}
	helper := dummy

	for l1 != nil || l2 != nil || carry != 0 {
		var l1Val, l2Val int
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		sum := l1Val + l2Val + carry
		carry = sum / 10
		helper.Next = &ListNode{sum % 10, nil}
		helper = helper.Next
	}
	return dummy.Next
}

type TreeNode struct {
	Val   int
	Next  *TreeNode
	Left  *TreeNode
	Right *TreeNode
}

type NodeSumPair struct {
	Node *TreeNode
	Sum  int
}

func PathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	queue := []*NodeSumPair{{root, root.Val}}

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.Node.Left != nil && curr.Node.Right != nil {
			if curr.Sum == targetSum {
				return true
			}
		} else {
			if curr.Node.Left != nil {
				queue = append(queue, &NodeSumPair{curr.Node.Left, curr.Sum + curr.Node.Left.Val})
			}
			if curr.Node.Right != nil {
				queue = append(queue, &NodeSumPair{curr.Node.Right, curr.Sum + curr.Node.Right.Val})
			}
		}
	}
	return false
}

func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var values []int
	stack := []TreeNode{*root}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		values = append(values, curr.Val)
		if curr.Right != nil {
			stack = append(stack, *curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, *curr.Left)
		}
	}

	return values
}

func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var values []int
	stack := []TreeNode{*root}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		values = append(values, curr.Val)
		if curr.Left != nil {
			stack = append(stack, *curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, *curr.Right)
		}
	}
	slices.Reverse(values)
	return values
}

func LLCycle(head *ListNode) bool {
	visited := make(map[*ListNode]bool)
	curr := head

	for curr != nil {
		if visited[curr] {
			return true
		}
		visited[curr] = true
		curr = curr.Next
	}
	return false
}

func LLIntersection(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	memory := make(map[*ListNode]bool)
	helperA := headA
	for helperA != nil {
		memory[helperA] = true
	}

	helperB := headB
	for helperB != nil {
		if _, exists := memory[helperB]; exists {
			return helperB
		}
		helperB = headB.Next
	}
	return nil
}

func LongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	memory := make(map[string]int, 0)
	lhs := 0
	result := 0

	for rhs := 0; rhs < len(s); rhs++ {
		if _, exists := memory[string(s[rhs])]; exists && memory[string(s[rhs])] >= lhs {
			lhs = memory[string(s[rhs])] + 1
		} else {
			result = max(result, rhs-lhs+1)
		}
		memory[string(s[rhs])] = rhs
	}
	return result
}

func MaxArea(heights []int) int {
	lhs := 0
	rhs := len(heights)
	if rhs == 0 {
		return 0
	}
	rhs -= 1

	maxArea := 0

	for lhs != rhs {
		base := rhs - lhs
		h := min(heights[lhs], heights[rhs])
		maxArea = max(maxArea, base*h)
		if heights[lhs] <= heights[rhs] {
			lhs += 1
		} else if heights[rhs] <= heights[lhs] {
			rhs -= 1
		} else {
			lhs += 1
		}
	}
	return maxArea
}
