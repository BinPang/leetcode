package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	t0 := reverseBetween(_arrayToListNode([]int{1, 2, 3, 4, 5}), 2, 4)
	printListNode(t0)

	t0 = reverseBetween(_arrayToListNode([]int{1, 2, 3, 4, 5}), 1, 1)
	printListNode(t0)
	t0 = reverseBetween(_arrayToListNode([]int{1, 2, 3, 4, 5}), 2, 2)
	printListNode(t0)
	t0 = reverseBetween(_arrayToListNode([]int{1, 2, 3, 4, 5}), 5, 5)
	printListNode(t0)

	t0 = reverseBetween(_arrayToListNode([]int{1, 2, 3, 4, 5}), 1, 2)
	printListNode(t0)

	t0 = reverseBetween(_arrayToListNode([]int{1, 2, 3, 4, 5}), 3, 5)
	printListNode(t0)
}

//Note: 1 ≤ m ≤ n ≤ length of list.
//Input: 1->2->3->4->5->NULL, m = 2, n = 4
//Output: 1->4->3->2->5->NULL
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	newHead := &ListNode{Next: head}
	var tmp *ListNode
	prev := newHead
	start := head
	for i := 1; i < m; i++ {
		start = start.Next
		prev = prev.Next
	}
	for i := m; i < n; i++ {
		tmp = start.Next
		start.Next = tmp.Next

		tmp.Next = prev.Next
		prev.Next = tmp
	}
	return newHead.Next
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
