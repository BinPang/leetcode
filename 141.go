package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	t1 := &ListNode{}
	t1.Next = t1
	println(hasCycle(t1))
	println(hasCycle(nil))
	println(hasCycle(&ListNode{}))
	println(hasCycle(&ListNode{Next: &ListNode{}}))
	println(hasCycle(&ListNode{Next: &ListNode{Next: &ListNode{}}}))
	println(hasCycle(&ListNode{Next: &ListNode{Next: &ListNode{Next: &ListNode{}}}}))
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p1 := head
	p2 := head
	for {
		if p2 == nil {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
		if p2 == nil {
			return false
		}
		p2 = p2.Next
		if p1 == p2 {
			return true
		}
	}
}
