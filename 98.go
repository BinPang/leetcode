package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	println("[1,null,1]:", isValidBST(&TreeNode{Val: 1, Right: &TreeNode{Val: 1}}))
	println("[2,1,3]:", isValidBST(&TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}))
	println("[5,1,4,null,null,3,6]", isValidBST(&TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right:&TreeNode{
				Val:6,
			},
		},
	}))
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isLeftOk(root.Left, root.Val, root.Val, true) && isRightOk(root.Right, root.Val, root.Val, true)
}

func isLeftOk(node *TreeNode, min, max int, extra bool) bool {
	if node == nil {
		return true
	}
	if extra {
		if node.Val >= max {
			return false
		}
	} else {
		if node.Val <= min || node.Val >= max {
			return false
		}
	}
	return isLeftOk(node.Left, min, node.Val, extra) && isRightOk(node.Right, node.Val, max, false)
}

func isRightOk(node *TreeNode, min, max int, extra bool) bool {
	if node == nil {
		return true
	}
	if extra {
		if node.Val <= max {
			return false
		}
	} else {
		if node.Val <= min || node.Val >= max {
			return false
		}
	}
	return isLeftOk(node.Left, min, node.Val, false) && isRightOk(node.Right, node.Val, max, extra)
}
