package main

func main() {
	t := levelOrderBottom(&TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	})
	for _, v := range t {
		for _, v1 := range v {
			print(v1, ",")
		}
		println("")
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	r := make([][]int, 0)
	reOrder(root, 0, &r)
	i:=0
	j:=len(r)-1
	for {
		if i >= j {
			break
		}
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return r
}

func reOrder(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}
	if len(*result) < level+1 {
		*result = append(*result, make([]int, 0))
		//*result = append([][]int{make([]int, 0)}, *result...)
	}
	if root.Left != nil {
		reOrder(root.Left, level+1, result)
	}
	(*result)[level] = append((*result)[level], root.Val)
	if root.Right != nil {
		reOrder(root.Right, level+1, result)
	}
}
