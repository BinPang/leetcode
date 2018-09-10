package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	r := _arrayToTreeNode([]interface{}{3, 9, 20, nil, nil, 15, 7})
	_printArrayArray(levelOrder(r))
}

func levelOrder(root *TreeNode) [][]int {
	r := make([][]int, 0)
	_levelOrder(root, &r, 0)
	return r
}

func _levelOrder(root *TreeNode, r *[][]int, level int) {
	if root == nil {
		return
	}
	if len(*r) == level {
		*r = append(*r, make([]int, 0))
	}
	(*r)[level] = append((*r)[level], root.Val)
	_levelOrder(root.Left, r, level+1)
	_levelOrder(root.Right, r, level+1)
}

func _arrayToTreeNode(a []interface{}) *TreeNode {
	t := make([]*TreeNode, len(a))
	for i, v := range a {
		t[i] = &TreeNode{}
		if i == 0 {
			if v == nil {
				return nil
			}
		} else {
			before := (i+1)/2 - 1
			lr := (i + 1) % 2
			if t[before] == nil {
				println("before is nil,i=", i, ",before=", before)
				panic("panic")
			}
			if v == nil {
				t[i] = nil
			}
			if lr == 0 {
				t[before].Left = t[i]
			} else {
				t[before].Right = t[i]
			}
		}
		if v != nil {
			t[i].Val = v.(int)
		}
	}
	return t[0]
}

func _printArrayArray(t [][]int) {
	for _, v := range t {
		_printArray(v)
	}
}
func _printArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}
