package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	if root.Left == nil {
		root.Left = root.Right
		root.Right = nil
		invertTree(root.Left)
	} else if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		invertTree(root.Right)
	} else {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}
