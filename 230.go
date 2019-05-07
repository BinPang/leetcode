package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t0 := &TreeNode{
		Val:3,
		Left:&TreeNode{
			Val:1,
			Right:&TreeNode{Val:2},
		},
		Right:&TreeNode{Val:4},
	}
	println(kthSmallest(t0, 1))
	println(kthSmallest(t0, 2))
	println(kthSmallest(t0, 3))
	println(kthSmallest(t0, 4))

	t1 := &TreeNode{
		Val:5,
		Left:&TreeNode{
			Val:3,
			Left:&TreeNode{
				Val:2,
				Left:&TreeNode{Val:1},
			},
			Right:&TreeNode{Val:4},
		},
		Right:&TreeNode{Val:6},
	}
	println(kthSmallest(t1, 3))
	println(kthSmallest(t1, 4))
}

func kthSmallest(root *TreeNode, k int) int {
	v, _ := _kthSmallest(root, k)

	return v
}

func _kthSmallest(root *TreeNode, k int) (int, bool) {
	if root == nil {
		return 0, false
	}
	v, ok := _kthSmallest(root.Left, k)
	if ok {
		return v, ok
	}
	v++
	if v == k {
		return root.Val, true
	}
	return _kthSmallest(root.Right, v)
}
