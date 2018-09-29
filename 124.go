package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*leetcode test case
[1,2,3]
[-2,20,20,null,null,15,7]
[-50,10,-40,3,10,5,0,null,null,null,null,10,-5]
[-1,-2,-3]
[]
---below is testing my wrong answer
[2,-1]
[-2,-1]
*/
func main() {
	t1 := _arrayToTreeNode([]interface{}{-2, -1, -3})
	println("input:[-2, -1, -3],need -1,output:", maxPathSum(t1))
	return

	t0 := _arrayToTreeNode([]interface{}{-10, 9, 20, nil, nil, 15, 7})
	println("input:[-10,9,20,null,null,15,7],need 42,output:", maxPathSum(t0))
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(_maxPathSum(root))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func _maxPathSum(root *TreeNode) (int, int) {
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	}
	if root.Left == nil || root.Right == nil {
		sum, path := 0, 0
		if root.Left == nil {
			sum, path = _maxPathSum(root.Right)
		} else {
			sum, path = _maxPathSum(root.Left)
		}
		return max(root.Val, max(sum, path+root.Val)), max(root.Val, path+root.Val)
	}
	sum0, path0 := _maxPathSum(root.Left)
	sum1, path1 := _maxPathSum(root.Right)
	newSum := max(path0+path1+root.Val, root.Val)
	if sum0 < sum1 {
		sum0 = sum1
	}
	if path0 < path1 {
		path0 = path1
	}
	return max(newSum, sum0), max(root.Val+path0, root.Val)
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
