package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return _constructMaximumBinaryTree(nums, 0, len(nums)-1)
}

func _constructMaximumBinaryTree(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	} else if start == end {
		return &TreeNode{Val: nums[start]}
	} else {
		m := end
		for i := start; i < end; i++ {
			if nums[i] > nums[m] {
				m = i
			}
		}
		return &TreeNode{
			Left:  _constructMaximumBinaryTree(nums, start, m-1),
			Val:   nums[m],
			Right: _constructMaximumBinaryTree(nums, m+1, end),
		}
	}
}
