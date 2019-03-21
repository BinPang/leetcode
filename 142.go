package main

type ListNode struct {
	Val  int
	Next *ListNode
}
/*
m :head to start
r :step of the loop
k :start to meet point
fast: m + t0 * r + k
slow: m + t1 * r + k
fast is two times of slow
m + t0 * r + k = 2*(m + t1 * r + k)
=> 0 = m + k + 2t1 * r - t0 * r
=> m + k = (t0 - 2t1) * r
t0 - 2t1 must be integer
so m + k must be some time(s) of circle
so meet point continue to go m step(s) must reach start point
but we don't know m exact, but we can move from head, m step(s) is the start point
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p1 := head
	p2 := head //p2 is fast
	for {
		if p2 == nil || p2.Next == nil {
			return nil
		}
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			break
		}
	}
	p1 = head
	for {
		if p1 == p2 {
			return p1
		} else {
			p1 = p1.Next
			p2 = p2.Next
		}
	}
}
