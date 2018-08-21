package main

func main() {
	ne := mergeTwoLists(&ListNode{1, &ListNode{7, &ListNode{9, nil}}}, &ListNode{2, &ListNode{3, &ListNode{4, nil}}})
	print("need:1, 2, 3, 4, 7, 9,return:")
	printListNode(ne)

	ne = mergeTwoLists(&ListNode{1, &ListNode{7, &ListNode{9, nil}}}, nil)
	print("need:1, 7, 9,return:")
	printListNode(ne)

	ne = mergeTwoLists(nil, &ListNode{1, &ListNode{7, &ListNode{9, nil}}})
	print("need:1, 7, 9,return:")
	printListNode(ne)

	ne = mergeTwoLists(&ListNode{2, &ListNode{3, &ListNode{4, nil}}}, &ListNode{1, &ListNode{7, &ListNode{9, nil}}})
	print("need:1, 2, 3, 4, 7, 9,return:")
	printListNode(ne)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(l *ListNode)  {
	for  {
		if l != nil {
			print(l.Val , ",")
			l = l.Next
		} else {
			print("\n")
			break
		}
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var rHead *ListNode
	var tp *ListNode
	if l1.Val > l2.Val {
		rHead = l2
		l2 = l2.Next
	} else {
		rHead = l1
		l1 = l1.Next
	}
	tp = rHead
	for {
		if l1 == nil {
			tp.Next = l2
			break
		}
		if l2 == nil {
			tp.Next = l1
			break
		}
		if l1.Val > l2.Val {
			tp.Next = l2
			tp = l2
			l2 = l2.Next
		} else {
			tp.Next = l1
			tp = l1
			l1 = l1.Next
		}

	}
	return rHead
}
