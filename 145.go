package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	println(fmt.Sprintf("%+v", postorderTraversal(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}})))
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	r := make([]int, 0)
	left := root.Left
	right := root.Right
	l := list.New()
	root.Left = nil
	root.Right = nil
	l.PushFront(root)
	if right != nil {
		l.PushFront(right)
	}
	if left != nil {
		l.PushFront(left)
	}
	for {
		if l.Len() == 0 {
			break
		}
		item := l.Front()
		node := item.Value.(*TreeNode)
		left = node.Left
		right = node.Right
		if node.Left == nil && node.Right == nil {
			r = append(r, node.Val)
			l.Remove(item)
		} else {
			node.Left = nil
			node.Right = nil
			if right != nil {
				l.PushFront(right)
			}
			if left != nil {
				l.PushFront(left)
			}
		}
	}
	return r
}
