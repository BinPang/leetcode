package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
Input: 2->1->3->5->6->4->7->NULL
Output: 2->3->6->7->1->5->4->NULL

Input: 1->2->3->4->5->NULL
Output: 1->3->5->2->4->NULL
*/

func main() {
	t0 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	r0 := oddEvenList(t0)
	println(r0)

	t1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	r1 := oddEvenList(t1)
	println(r1)

	t2 := &ListNode{Val: 2,
		Next: &ListNode{Val: 1,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 5,
					Next: &ListNode{Val: 6,
						Next: &ListNode{Val: 4,
							Next: &ListNode{Val: 7}}}}}}}
	r2 := oddEvenList(t2)
	println(r2)
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	likeZero := true //0, 2, 4, 6, 8...
	p1 := head
	p2 := head.Next
	th := p2
	tmp := head.Next.Next
	for tmp != nil {
		if likeZero {
			p1.Next = tmp
			p1 = tmp
		} else {
			p2.Next = tmp
			p2 = tmp
		}
		tmp = tmp.Next
		likeZero = !likeZero
	}
	p1.Next = th
	p2.Next = nil
	return head
}
