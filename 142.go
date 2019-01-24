package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p1 := head
	p2 := head
	for {
		if p2 == nil {
			return nil
		}
		p1 = p1.Next
		p2 = p2.Next
		if p2 == nil {
			return nil
		}
		p2 = p2.Next
		if p1 == p2 {
			return p1
		}
	}
}
