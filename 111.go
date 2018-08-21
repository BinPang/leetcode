package main

func main() {
	println("need 2,return", minDepth(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}))
	println("need 2,return", minDepth(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}))
	println("need 0,return", minDepth(nil))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m := 0
	_minDepth(root, 1, &m)
	return m

}

func _minDepth(root *TreeNode, depth int, m *int) {
	if root.Left == nil && root.Right == nil {
		if *m == 0 || *m > depth {
			*m = depth
		}
		return
	}
	if root.Left != nil {
		_minDepth(root.Left, depth+1, m)
	}
	if root.Right != nil {
		_minDepth(root.Right, depth+1, m)

	}
}
