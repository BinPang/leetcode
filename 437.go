package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return _pathSum(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func _pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	if root.Val == sum {
		return 1 + _pathSum(root.Left, sum-root.Val) + _pathSum(root.Right, sum-root.Val)
	} else {
		return _pathSum(root.Left, sum-root.Val) + _pathSum(root.Right, sum-root.Val)
	}
}
