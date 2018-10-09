package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t0 := _arrayToTreeNodeLeetCode([]interface{}{3, 1, 4, nil, nil, 2}) //test cast next exchange
	inOrder(t0)
	print("need:{1,2,3,4,},result:")
	recoverTree(t0)
	inOrder(t0)
	println("")

	t1 := _arrayToTreeNodeLeetCode([]interface{}{10, 20, 7, 6, 8, 15, 30}) //test case two diff exchange
	inOrder(t1)
	print("need:{6,7,8,10,15,20,30,},result:")
	recoverTree(t1)
	inOrder(t1)
	println("")
}

func recoverTree(root *TreeNode) {
	var m1, m2 TreeNode
	var pre TreeNode
	_recoverTree(root, &m1, &m2, &pre)
	m1.Left.Val, m2.Left.Val = m2.Left.Val, m1.Left.Val
}

func _recoverTree(root *TreeNode, m1, m2, pre *TreeNode) {
	//defer func() {
	//	println("end:", root, root.Val, m1.Val, m2.Val, pre.Val, m1.Left, m2.Left, pre.Left)
	//}()
	//println("start:", root, root.Val, m1.Val, m2.Val, pre.Val, m1.Left, m2.Left, pre.Left)
	if root.Left == nil && root.Right == nil {
		if pre.Left == nil {
			pre.Left = root
		} else if pre.Left.Val > root.Val {
			if m2.Left == nil {
				m1.Left = pre.Left
				m2.Left = root
			} else {
				m2.Left = root
			}
		}
		pre.Left = root
		return
	}
	if root.Left != nil {
		_recoverTree(root.Left, m1, m2, pre)
	}
	if pre.Left != nil {
		if pre.Left.Val > root.Val {
			if m2.Left == nil {
				m1.Left = pre.Left
				m2.Left = root
			} else {
				m2.Left = root
			}
		}
	}
	pre.Left = root
	if root.Right != nil {
		_recoverTree(root.Right, m1, m2, pre)
	}
	return
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

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	print(node.Val, ",")
	inOrder(node.Right)
}
