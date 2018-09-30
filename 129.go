package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	println("input:[4,9,0,5,1],need:1026,result:", sumNumbers(_arrayToTreeNodeLeetCode([]interface{}{4, 9, 0, 5, 1})))
}

func sumNumbers(root *TreeNode) int {
	return _sumNumbers(root, 0)
}

func _sumNumbers(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}
	if root.Left == nil && root.Right == nil {
		return sum*10 + root.Val
	}
	if root.Left == nil {
		return _sumNumbers(root.Right, sum*10+root.Val)
	}
	if root.Right == nil {
		return _sumNumbers(root.Left, sum*10+root.Val)
	}
	return _sumNumbers(root.Left, sum*10+root.Val) + _sumNumbers(root.Right, sum*10+root.Val)
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
