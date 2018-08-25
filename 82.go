package main

func main() {
	println("input:1->1->2, need:2")
	printListNode(deleteDuplicates(_arrayToListNode([]int{1, 1, 2})))

	println("input:1->1, need:")
	printListNode(deleteDuplicates(_arrayToListNode([]int{1, 1})))

	println("input:1->2->2, need:1")
	printListNode(deleteDuplicates(_arrayToListNode([]int{1, 2, 2})))

	println("input:1, need:1")
	printListNode(deleteDuplicates(_arrayToListNode([]int{1})))

	println("input:1->2->3->3->4->4->5, need:1->2->5")
	printListNode(deleteDuplicates(_arrayToListNode([]int{1, 2,3,3,4,4,5})))

	println("input:1->1->1->2->3, need:2->3")
	printListNode(deleteDuplicates(_arrayToListNode([]int{1, 1, 1, 2, 3})))

	println("input:1->1->1, need:")
	printListNode(deleteDuplicates(_arrayToListNode([]int{1, 1, 1})))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tmp := head
	next := head
	delSelf := false
	var newHead *ListNode
	var prev *ListNode
	for {
		if tmp == nil {
			break
		}
		for {
			next = tmp.Next
			if next == nil {
				if newHead == nil {
					newHead = tmp
				}
				if delSelf {
					if prev == nil {
						newHead = nil
					} else {
						prev.Next = nil
					}
				}
				tmp = nil
				break
			}
			if next.Val != tmp.Val {
				if delSelf {
					if prev != nil {
						prev.Next = next
					}
				} else {
					prev = tmp
					if newHead == nil {
						newHead = tmp
					}
				}
				delSelf = false
				tmp = next
			} else {
				delSelf = true
				tmp.Next = next.Next
			}
		}
	}
	return newHead
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
