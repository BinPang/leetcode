package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	vLeft := 0
	vRight := 0
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			vLeft = root.Left.Val
		} else {
			vLeft = sumOfLeftLeaves(root.Left)
		}
	}
	vRight = sumOfLeftLeaves(root.Right)
	return vLeft + vRight
}
