package main

func main() {

	t4 := reverseKGroup(_arrayToListNode([]int{1}), 2)
	println("input:1, 2")
	printListNode(t4)

	t3 := reverseKGroup(_arrayToListNode([]int{1, 2, 3, 4, 5}), 5)
	println("input:1->2->3->4->5, 2")
	printListNode(t3)

	t2 := reverseKGroup(_arrayToListNode([]int{1, 2, 3, 4, 5}), 2)
	println("input:1->2->3->4->5, 2")
	printListNode(t2)

	t1 := reverseKGroup(_arrayToListNode([]int{1, 2, 3, 4, 5}), 4)
	println("input:1->2->3->4->5, 4")
	printListNode(t1)



}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	var r *ListNode

	start := head
	end := head
	tmpK := k
	tmpStart := head
	var prev *ListNode
	for {
		tmpK = k
		for {
			if end == nil {
				break
			}
			tmpK --
			end = end.Next
			if tmpK == 1 || end == nil {
				break
			}
		}
		if end == nil {
			break
		}
		if r == nil {
			r = end
		}
		if prev != nil {
			prev.Next = end
		}
		for {
			if start == end {
				break
			}
			tmpStart = start.Next
			start.Next = end.Next
			end.Next = start
			start = tmpStart
		}
		tmpK = k
		for  {
			tmpK--
			end = end.Next
			if tmpK == 1 {
				break
			}
		}
		prev = end
		start = prev.Next
		end = prev.Next
	}
	if r == nil {
		return head
	}
	return r
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
func _arrayToListNode(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	head := &ListNode{Val: a[0]}
	prev := head
	for i, v := range a {
		if i == 0 {
			continue
		}
		prev.Next = &ListNode{Val: v}
		prev = prev.Next
	}
	return head
}
