package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	r := make([]int, 0)
	_largestValues(root, &r, 0)
	return r
}

func _largestValues(node *TreeNode, r *[]int, level int) {
	if node == nil {
		return
	}
	if len(*r) == level {
		*r = append(*r, node.Val)
	} else {
		if (*r)[level] < node.Val {
			(*r)[level] = node.Val
		}
	}
	_largestValues(node.Left, r, level+1)
	_largestValues(node.Right, r, level+1)
}
