package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t := sortedListToBST(_arrayToListNode([]int{-10,-3,0,5,9}))
	inOrder(t)
	println("")
}

func sortedListToBST(head *ListNode) *TreeNode {
	return _sortedListToBST(head, nil)
}

func _sortedListToBST(start, end *ListNode) *TreeNode {
	if start == end {
		return nil
	}
	fast := start
	slow := start
	for {
		if fast == end || fast.Next == end {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return &TreeNode{
		Val:   slow.Val,
		Left:  _sortedListToBST(start, slow),
		Right: _sortedListToBST(slow.Next, end),
	}
}


func _arrayToListNode(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	head := &ListNode{Val:a[0]}
	prev := head
	for i, v := range a{
		if i==0 {
			continue
		}
		prev.Next = &ListNode{Val:v}
		prev = prev.Next
	}
	return head
}


func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	print(node.Val, ",")
	inOrder(node.Left)
	inOrder(node.Right)
}
