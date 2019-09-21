package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t0 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	printMatrix(printTree(t0))

	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 3},
	}
	printMatrix(printTree(t1))
}

func printMatrix(data [][]string) {
	for _, v := range data {
		for _, vv := range v {
			if vv == "" {
				print("\"\"\t")
			} else {
				print(fmt.Sprintf("\"%s\"\t", vv))
			}
		}
		println()
	}
	println()
}

func printTree(root *TreeNode) [][]string {
	if root == nil {
		return nil
	}
	h := _findHeight(root)
	var i uint
	r := make([][]string, h)
	tmp := make(map[uint]*TreeNode)
	tmp[1<<(h-1)-1] = root
	for i = 0; i < h; i++ {
		r[i] = make([]string, 1<<h-1)
		next := make(map[uint]*TreeNode)
		for k, v := range tmp {
			r[i][k] = strconv.Itoa(v.Val)
			if v.Left != nil {
				next[k-(1<<(h-i-2))] = v.Left
			}
			if v.Right != nil {
				next[k+(1<<(h-i-2))] = v.Right
			}
		}
		tmp = next
	}
	return r
}

func _findHeight(root *TreeNode) uint {
	if root == nil {
		return 0
	}
	return max(_findHeight(root.Left), _findHeight(root.Right)) + 1
}
func max(a, b uint) uint {
	if a > b {
		return a
	} else {
		return b
	}
}
