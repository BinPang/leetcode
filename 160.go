package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tmpLenHeadA := headA
	lenHeadA := 0
	for {
		if tmpLenHeadA == nil {
			break
		}
		lenHeadA ++
		tmpLenHeadA = tmpLenHeadA.Next
	}

	tmp := make(map[*ListNode]bool, lenHeadA)
	for {
		if headA == nil {
			break
		}
		tmp[headA] = true
		headA = headA.Next
	}
	for {
		if headB == nil {
			break
		}
		if tmp[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
