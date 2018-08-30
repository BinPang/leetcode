package main

func main() {

	l1:= rotateRight(_arrayToListNode([]int{1, 2, 3, 4, 5}), 7)
	printListNode(l1)

	l0 := rotateRight(_arrayToListNode([]int{1, 2, 3, 4, 5}), 2)
	printListNode(l0)

	l2 := rotateRight(_arrayToListNode([]int{1, 2, 3, 4, 5}), 5)
	printListNode(l2)

	l3 := rotateRight(_arrayToListNode([]int{1, 2, 3, 4, 5}), 4)
	printListNode(l3)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	l := 0
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head
	for tmp != nil {
		l++
		tmp = tmp.Next
	}
	k = k % l
	if k <= 0 {
		return head
	}
	step := l - k
	r := head
	for {
		if step == 1 {
			tmp := r
			r = r.Next
			tmp.Next = nil
			break
		} else {
			r = r.Next
		}
		step--
	}
	tmp = r
	for {
		if tmp.Next == nil {
			tmp.Next = head
			break
		}
		tmp = tmp.Next
	}
	return r
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
