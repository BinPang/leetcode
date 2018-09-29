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
[-10, 9, 20, nil, nil, 15, 7]
*/
func main() {
	t4 := _arrayToTreeNodeLeetCode([]interface{}{-2,-1})
	println("input:[-2,-1],need -1,output:", maxPathSum(t4))

	t3 := _arrayToTreeNodeLeetCode([]interface{}{2,-1})
	println("input:[2,-1],need 2,output:", maxPathSum(t3))

	t2 := _arrayToTreeNodeLeetCode([]interface{}{-1, nil, 9, -6, 3, nil, nil, nil, -2})
	println("input:[-1,nil,9,-6,3,nil,nil,nil,-2],need 12,output:", maxPathSum(t2))

	t1 := _arrayToTreeNodeLeetCode([]interface{}{-2, -1, -3})
	println("input:[-2, -1, -3],need -1,output:", maxPathSum(t1))

	t0 := _arrayToTreeNodeLeetCode([]interface{}{-10, 9, 20, nil, nil, 15, 7})
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
	//newSum := max(path0+path1+root.Val, root.Val)
	newSum := max(root.Val, max(path0+path1+root.Val, max(path0+root.Val, path1+root.Val)))
	if sum0 < sum1 {
		sum0 = sum1
	}
	if path0 < path1 {
		path0 = path1
	}
	return max(newSum, sum0), max(root.Val+path0, root.Val)
}

func _arrayToTreeNodeLeetCode(a []interface{}) *TreeNode {
	l := len(a)
	if l == 0 || a[0] == nil {
		return nil
	}
	r := &TreeNode{Val: a[0].(int)}
	lastLevel := make([]*TreeNode, 0) //last level *TreeNode
	lastLevel = append(lastLevel, r)
	lastIndex := 0
	nowLevel := make([]*TreeNode, 0)
	i := 1
	for {
		if i >= l {
			break
		}
		if len(lastLevel) == 0 {
			break
		}
		if a[i] == nil {
			if lastIndex%2 == 0 {
				lastLevel[lastIndex/2].Left = nil
			} else {
				lastLevel[lastIndex/2].Right = nil
			}
		} else {
			tmp := &TreeNode{Val: a[i].(int)}
			nowLevel = append(nowLevel, tmp)
			if lastIndex%2 == 0 {
				lastLevel[lastIndex/2].Left = tmp
			} else {
				lastLevel[lastIndex/2].Right = tmp
			}
		}
		lastIndex ++
		i++
		if lastIndex == len(lastLevel)*2 {
			lastIndex = 0
			lastLevel = nowLevel
			nowLevel = make([]*TreeNode, 0)
		}
	}
	return r
}
