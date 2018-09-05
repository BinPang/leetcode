package main

type ListNode struct {
	Val  int
	Next *ListNode
}


//Input: 1->2->3->4->5->NULL, m = 2, n = 4
//Output: 1->4->3->2->5->NULL
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	addedHead := &ListNode{Next:head}
	prev := addedHead
	start := head
	for i:=1;i<m ;i++{
		start = start.Next
		prev = prev.Next
	}
}
