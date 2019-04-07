package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 6},
		},
	}
	println(countNodes(t))
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := leafLen(root.Left, true)
	right := leafLen(root.Right, false)
	if left == right {
		return 2<<left - 1
	} else {
		return countNodes(root.Left) + countNodes(root.Right) + 1
	}
}

func leafLen(root *TreeNode, left bool) uint {
	var r uint = 0
	for root != nil {
		r++
		if left {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return r
}
