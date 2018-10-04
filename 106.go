package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t1 := buildTree([]int{17, 15, 12, 9, 6, 3}, []int{17, 15, 12, 9, 6, 3})
	inOrder(t1)
	println(":end:3,6,9,12,15,17")
	return

	t := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	inOrder(t)
	println(":end:3, 9, 20, 15, 7")
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	return _buildTree(inorder, postorder, 0, len(postorder)-1, 0, len(inorder)-1)
}

func _buildTree(inorder []int, postorder []int, pStart, pEnd, iStart, iEnd int) *TreeNode {
	//println("input:", pStart, pEnd, iStart, iEnd)
	if pStart > pEnd {
		return nil
	}
	node := &TreeNode{Val: postorder[pEnd]}
	if pStart == pEnd {
		return node
	}
	nowpEnd := pStart
	nowiEnd := iStart
	for {
		if postorder[pEnd] == inorder[nowiEnd] {
			break
		}
		nowiEnd += 1
		nowpEnd += 1
	}
	//println("left:", pStart+1, nowpEnd, iStart, nowiEnd-1)
	node.Left = _buildTree(inorder, postorder, pStart, nowpEnd-1, iStart, nowiEnd-1)

	//println("right:", nowpEnd+1, pEnd, nowiEnd+1, iEnd)
	node.Right = _buildTree(inorder, postorder, nowpEnd, pEnd-1, nowiEnd+1, iEnd)

	return node
}

//it's pre order
func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	print(node.Val, ",")
	inOrder(node.Left)
	inOrder(node.Right)
}
