package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	v, _ := _findBottomLeftValue(root, 0)
	return v
}

func _findBottomLeftValue(root *TreeNode, level int) (int, int) {
	if root.Left == nil && root.Right == nil {
		return root.Val, level + 1
	}
	if root.Left == nil {
		return _findBottomLeftValue(root.Right, level+1)
	} else if root.Right == nil {
		return _findBottomLeftValue(root.Left, level+1)
	}
	lv, ll := _findBottomLeftValue(root.Left, level+1)
	rv, rl := _findBottomLeftValue(root.Right, level+1)
	if rl > ll {
		return rv, rl
	} else {
		return lv, ll
	}
}
