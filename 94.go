package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//easy, no test case,only use leetcode run code method

func inorderTraversal(root *TreeNode) []int {
	r := make([]int, 0)
	_inorderTraversal(root, &r)
	return r
}

func _inorderTraversal(root *TreeNode, r *[]int)  {
	if root == nil {
		return
	}
	_inorderTraversal(root.Left, r)
	*r = append(*r, root.Val)
	_inorderTraversal(root.Right, r)
}



