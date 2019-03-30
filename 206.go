package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fake := &ListNode{Next:head}
	prev := head
	now := head.Next
	for {
		if now == nil {
			break
		}
		prev.Next = now.Next
		now.Next = fake.Next
		fake.Next = now

		now = prev.Next
	}
	return fake.Next
}
