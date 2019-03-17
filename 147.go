package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	t3 := insertionSortList(_arrayToListNode([]int{6, 5, 7, 8}))
	printListNode(t3)

	t2 := insertionSortList(_arrayToListNode([]int{5, 7, 6, 8}))
	printListNode(t2)

	t1 := insertionSortList(_arrayToListNode([]int{5, 6, 7, 8}))
	printListNode(t1)
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//add a fake head
	faker := &ListNode{Next: head}
	pointNow := head.Next
	prev := head
	loopPrev := prev
	loopPoint := pointNow

	for {
		if pointNow == nil {
			break
		}
		prev.Next = pointNow.Next

		loopPrev = faker
		loopPoint = faker.Next
		for {
			if pointNow.Val < loopPoint.Val {
				pointNow.Next = loopPoint
				loopPrev.Next = pointNow
				break
			} else {
				if loopPoint.Next == pointNow.Next {
					prev.Next = pointNow
					prev = prev.Next
					break
				} else {
					loopPrev = loopPrev.Next
					loopPoint = loopPoint.Next
				}
			}
		}
		pointNow = prev.Next
	}

	return faker.Next
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
