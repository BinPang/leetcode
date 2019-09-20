package main

func averageOfLevels(root *TreeNode) []float64 {
	r := make([]float64, 0)
	if root == nil {
		return r
	}
	node := []*TreeNode{root}
	for len(node) > 0 {
		tmp := make([]*TreeNode, 0)
		var tmpSum float64 = 0
		var tmpTotal float64 = 0
		for _, v := range node {
			tmpTotal += 1
			tmpSum += float64(v.Val)
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}
		node = tmp
		r = append(r, tmpSum/tmpTotal)
	}
	return r
}
