package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	tl := 0
	tr := 0
	if root.Left != nil {
		tl = maxDepth(root.Left)
	}
	if root.Right != nil {
		tr = maxDepth(root.Right)
	}
	if tl < tr {
		tl = tr
	}
	return tl + 1
}
