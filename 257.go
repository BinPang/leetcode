package main

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	r := make([]string, 0)
	if root == nil {
		return r
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	} else if root.Left == nil {
		for _, v := range binaryTreePaths(root.Right) {
			r = append(r, strconv.Itoa(root.Val)+"->"+v)
		}
	} else if root.Right == nil {
		for _, v := range binaryTreePaths(root.Left) {
			r = append(r, strconv.Itoa(root.Val)+"->"+v)
		}
	} else {
		for _, v := range binaryTreePaths(root.Left) {
			r = append(r, strconv.Itoa(root.Val)+"->"+v)
		}
		for _, v := range binaryTreePaths(root.Right) {
			r = append(r, strconv.Itoa(root.Val)+"->"+v)
		}
	}

	return r
}
