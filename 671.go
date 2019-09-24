package main

func findSecondMinimumValue(root *TreeNode) int {
	return _findSecondMinimumValue(root, root.Val)
}

func _findSecondMinimumValue(root *TreeNode, v int) int {
	if root == nil {
		return -1
	}
	if root.Val != v {
		return root.Val
	}
	l := _findSecondMinimumValue(root.Left, v)
	r := _findSecondMinimumValue(root.Right, v)
	if l == -1 {
		return r
	}
	if r == -1 {
		return l
	}
	if l > r {
		return r
	} else {
		return l
	}
}
