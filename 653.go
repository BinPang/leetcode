package main

func findTarget(root *TreeNode, k int) bool {
	tmp := make(map[int]struct{}, 0)
	return _findTarget(root, k, &tmp)
}

func _findTarget(root *TreeNode, k int, tmp *map[int]struct{}) bool {
	if root == nil {
		return false
	}
	if _, ok := (*tmp)[root.Val]; ok {
		return true
	}
	(*tmp)[k-root.Val] = struct{}{}

	return _findTarget(root.Left, k, tmp) || _findTarget(root.Right, k, tmp)
}
