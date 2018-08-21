package main

func main() {
	t := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	inOrder(t)
	println(":end:3, 9, 20, 15, 7")

	t = buildTree([]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{8, 7, 6, 5, 4, 3, 2, 1})
	inOrder(t)
	println(":end:1, 2, 3, 4, 5, 6, 7, 8")


	t = buildTree([]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8})
	inOrder(t)
	println(":end:1, 2, 3, 4, 5, 6, 7, 8")

	t = buildTree([]int{1,2,4,8,9,5,10,11,3,6,12,13,7,14,15}, []int{8,4,9,2,10,5,11,1,12,6,13,3,14,7,15})
	inOrder(t)
	println(":end:1,2,4,8,9,5,10,11,3,6,12,13,7,14,15")

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	return buildTreeNew(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}
func buildTreeNew(preorder []int, inorder []int, pStart, pEnd, iStart, iEnd int) *TreeNode {
	//println("input:", pStart, pEnd, iStart, iEnd)
	if pStart > pEnd {
		return nil
	}
	node := &TreeNode{Val: preorder[pStart]}
	if pStart == pEnd {
		return node
	}
	nowpEnd := pStart
	nowiEnd := iStart
	for {
		if preorder[pStart] == inorder[nowiEnd] {
			break
		}
		nowiEnd += 1
		nowpEnd += 1
	}
	//println("left:", pStart+1, nowpEnd, iStart, nowiEnd-1)
	node.Left = buildTreeNew(preorder, inorder, pStart+1, nowpEnd, iStart, nowiEnd-1)

	//println("right:", nowpEnd+1, pEnd, nowiEnd+1, iEnd)
	node.Right = buildTreeNew(preorder, inorder, nowpEnd+1, pEnd, nowiEnd+1, iEnd)

	return node
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	print(node.Val, ",")
	inOrder(node.Left)
	inOrder(node.Right)
}
