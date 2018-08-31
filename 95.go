package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t := generateTrees(3)
	println(len(t))
	for _, v := range t {
		inOrder(v)
		println("end")
	}
}

func generateTrees(n int) []*TreeNode {
	return _generateTrees(1, n)
}

func _generateTrees(start, end int) []*TreeNode {
	if start > end {
		return nil
	} else if start == end {
		return []*TreeNode{{Val: start}}
	}
	r := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		left := _generateTrees(start, i-1)
		right := _generateTrees(i+1, end)
		//println(i, len(left), len(right))
		if len(left) == 0 {
			for _, v1 := range right {
				r = append(r, &TreeNode{Val: i, Right: v1})
			}
		} else if len(right) == 0 {
			for _, v := range left {
				r = append(r, &TreeNode{Val: i, Left: v})
			}
		} else {
			for _, v0 := range left {
				for _, v1 := range right {
					r = append(r, &TreeNode{Val: i, Left: v0, Right: v1})
				}
			}
		}
	}
	//println("___", start, end, len(r))
	return r
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	print(node.Val, ",")
	inOrder(node.Left)
	inOrder(node.Right)
}
