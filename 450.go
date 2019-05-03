package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	origin := &TreeNode{
		Val: 50,
		Left: &TreeNode{
			Val:  30,
			Left: &TreeNode{Val: 20},
			Right: &TreeNode{
				Val:  40,
				Left: &TreeNode{Val: 35},
			},
		},
		Right: &TreeNode{
			Val:   60,
			Right: &TreeNode{Val: 70},
		},
	}
	del := deleteNode(origin, 30)
	println("end", del.Val)
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		return rotate(root)
	}
	prev := root
	now := root
	left := true
	for {
		if now == nil {
			break
		}
		if now.Val == key {
			if now.Left == nil {
				if left {
					prev.Left = now.Right
				} else {
					prev.Right = now.Right
				}
			} else if now.Right == nil {
				if left {
					prev.Left = now.Left
				} else {
					prev.Right = now.Left
				}
			} else {
				tmp := rotate(now)
				if left {
					prev.Left = tmp
				} else {
					prev.Right = tmp
				}
			}
			break
		} else if now.Val > key {
			prev = now
			now = now.Left
			left = true
		} else {
			prev = now
			now = now.Right
			left = false
		}
	}
	return root
}

func rotate(node *TreeNode) *TreeNode {
	tmp := node.Left.Right
	node.Left.Right = node.Right
	start := node.Right

	for {
		if start.Left == nil {
			start.Left = tmp
			break
		}
		start = start.Left
	}
	return node.Left
}
