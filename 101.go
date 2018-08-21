package main

func main() {
	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}
	println("true,", isSymmetric(t))
	t.Left.Left=nil
	println("false", isSymmetric(t))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return _isSymmetric(root.Left, root.Right)
}

func _isSymmetric(root *TreeNode, sub *TreeNode) bool {
	if root == nil && sub == nil {
		return true
	}
	if root == nil || sub == nil {
		return false
	}
	if root.Val != sub.Val {
		return false
	}
	l := _isSymmetric(root.Left, sub.Right)
	r := _isSymmetric(root.Right, sub.Left)
	return l && r
}
