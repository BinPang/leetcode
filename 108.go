package main

func main() {
	t := sortedArrayToBST([]int{-10,-3,0,5,9})
	inOrder(t)
	println("")

	t = sortedArrayToBST([]int{1, 2,3,4,5,6,7,8,9,10,11,12})
	inOrder(t)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return &TreeNode{Val: nums[0]}
	}
	if l == 2 {
		return &TreeNode{Val: nums[1], Left: &TreeNode{Val: nums[0]}}
	}
	if l == 3 {
		return &TreeNode{Val: nums[1], Left: &TreeNode{Val: nums[0]}, Right: &TreeNode{Val: nums[2]}}
	}
	i := l / 2
	return &TreeNode{
		Val:nums[i],
		Left:sortedArrayToBST(nums[0:i]),
		Right:sortedArrayToBST(nums[i+1:]),
	}
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	print(node.Val, ",")
	inOrder(node.Left)
	inOrder(node.Right)
}
