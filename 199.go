package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	rightSideView(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 4},
		},
	})
}
func rightSideView(root *TreeNode) []int {
	r := make([]int, 0)
	_rightSideView(root, 0, &r)
	return r
}

func _rightSideView(root *TreeNode, level int, r *[]int) {
	if root == nil {
		return
	}
	if level == len(*r) {
		*r = append(*r, 0)
	}
	_rightSideView(root.Left, level+1, r)
	(*r)[level] = root.Val
	_rightSideView(root.Right, level+1, r)
}
