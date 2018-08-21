package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	println(hasPathSum(&TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, 5))
	t := &TreeNode{}
	println(hasPathSum(t.Left, 0))
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if sum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left != nil && root.Right != nil {
		return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
	}
	if root.Left != nil {
		return hasPathSum(root.Left, sum-root.Val)
	}
	if root.Right != nil {
		return hasPathSum(root.Right, sum-root.Val)
	}
	return false
}
