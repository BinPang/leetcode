package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t := &TreeNode{
		Val:5,
		Left: &TreeNode{
			Val:4,
			Left:&TreeNode{
				Val:11,
				Left:&TreeNode{
					Val:7,
				},
				Right:&TreeNode{
					Val:2,
				},
			},
		},
		Right:&TreeNode{
			Val:8,
			Left:&TreeNode{
				Val:13,
			},
			Right:&TreeNode{
				Val:4,
				Left:&TreeNode{
					Val:5,
				},
				Right:&TreeNode{
					Val:1,
				},
			},
		},
	}
	tv := pathSum(t, 22)
	_printArrayArray(tv)

	tv = pathSum(nil, 22)
	println(tv)
	_printArrayArray(tv)
}

func _printArrayArray(t [][]int) {
	for _, v := range t{
		_printArray(v)
	}
}
func _printArray(t []int)  {
	for _, v := range t{
		print(v, ",")
	}
	println("")
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return [][]int{{root.Val}}
		} else {
			return nil
		}
	}
	var left [][]int
	var right [][]int
	if root.Left != nil {
		left = pathSum(root.Left, sum-root.Val)
	}
	if root.Right != nil {
		right = pathSum(root.Right, sum-root.Val)
	}
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		left, right = right, left
	}
	for i, v := range left {
		left[i] = append([]int{root.Val}, v...)
	}
	if right != nil {
		for _, v := range right {
			left = append(left, append([]int{root.Val}, v...))
		}
	}
	return left
}
