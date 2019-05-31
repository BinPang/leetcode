package main


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	if t1.Left == nil {
		t1.Left = t2.Left
	} else {
		mergeTrees(t1.Left, t2.Left)
	}
	if t1.Right == nil {
		t1.Right = t2.Right
	} else {
		mergeTrees(t1.Right, t2.Right)
	}
	return t1
}
