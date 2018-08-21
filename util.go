package main

func main() {
	inOrder(_arrayToTreeNode([]interface{}{3, 1, nil, nil, 2}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
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

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	print(node.Val, ",")
	inOrder(node.Left)
	inOrder(node.Right)
}

func printListNode(l *ListNode) {
	for {
		if l != nil {
			print(l.Val, ",")
			l = l.Next
		} else {
			print("\n")
			break
		}
	}
}
