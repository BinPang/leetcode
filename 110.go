package main

func main() {
	println("need false,result:", isBalanced(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}))
	//return
	t1 := isBalanced(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 2},
	})
	println("need false,result:", t1)
	//return
	t := isBalanced(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	})
	println("need true,result:", t)
	println("need true,result:", isBalanced(nil))
	println("need false,result:", isBalanced(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, b := check(root)
	return b
}

func check(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	left, leftBool := check(root.Left)
	right, rightBool := check(root.Right)
	//println(left, leftBool, right, rightBool)
	if !leftBool || !rightBool {
		return 0, false
	}
	if left > right {
		left, right = right, left
	}
	if right-left > 1 {
		return 0, false
	}
	return right + 1, true

}
