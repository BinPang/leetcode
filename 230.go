package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t0 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 4},
	}
	println(kthSmallest(t0, 1))
	println(kthSmallest(t0, 2))
	println(kthSmallest(t0, 3))
	println(kthSmallest(t0, 4))

	t1 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 6},
	}
	println(kthSmallest(t1, 3))
	println(kthSmallest(t1, 4))
}

func kthSmallest(root *TreeNode, k int) int {
	exist := 0
	return _kthSmallest(root, k, &exist)
}

func _kthSmallest(root *TreeNode, k int, exist *int) int {
	if root == nil {
		return 0
	}
	v := _kthSmallest(root.Left, k, exist)
	if *exist == k {
		return v
	}
	*exist ++
	if *exist == k {
		return root.Val
	}
	return _kthSmallest(root.Right, k, exist)
}
