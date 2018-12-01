package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	printListNode(partition(_arrayToListNode([]int{1, 4, 3, 2, 5, 2}), 3))
	printListNode(partition(_arrayToListNode([]int{1, 4, 3, 2, 5, 2}), 6))
	printListNode(partition(_arrayToListNode([]int{1, 4, 3, 2, 5, 2}), 0))
	printListNode(partition(_arrayToListNode([]int{1, 1, 3, 2, 5, 2}), 1))
}

//Input: head = 1->4->3->2->5->2, x = 3
//Output: 1->2->2->4->3->5
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fakerHead := &ListNode{Next: head}
	start := head
	run := head.Next
	prev := head
	for {
		if run == nil {
			break
		}
		if run.Val >= x {
			break
		} else {
			start = start.Next
			run = run.Next
			prev = prev.Next
		}
	}
	for {
		if run == nil {
			break
		}
		println(start, start.Val, run, run.Val, prev, prev.Val)
		if run.Val >= x {
			run = run.Next
			prev = prev.Next
		} else {
			prev.Next = run.Next
			run.Next = start.Next
			start.Next = run

			run = run.Next
		}
	}

	return fakerHead.Next
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
