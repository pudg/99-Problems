package thirty

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}
	helper := dummy

	for l1 != nil || l2 != nil || carry != 0 {
		var l1Val, l2Val int
		if l1 != nil {
			l1Val = l1.Val
		} else {
			l1Val = 0
		}

		if l2 != nil {
			l2Val = l2.Val
		} else {
			l2Val = 0
		}

		currSum := l1Val + l2Val + carry
		val := currSum % 10
		carry = currSum / 10
		helper.Next = &ListNode{val, nil}
		helper = helper.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return dummy.Next
}
