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
