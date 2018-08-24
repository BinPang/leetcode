package main

func main() {
	println("input:1->1->2, need:1->2")
	printListNode(deleteDuplicates(_arrayToListNode([]int{1, 1, 2})))

	println("input:1->1->2->3->3, need:1->2->3")
	printListNode(deleteDuplicates(_arrayToListNode([]int{1, 1, 2, 3, 3})))

	println("input:1, need:1")
	printListNode(deleteDuplicates(_arrayToListNode([]int{1})))
}

type ListNode struct {
	Val  int
	Next *ListNode
}
func deleteDuplicates(head *ListNode) *ListNode {
	tmp := head
	next := head
	for {
		if tmp == nil {
			break
		}
		for {
			next = tmp.Next
			if next == nil {
				tmp = nil
				break
			}
			if next.Val != tmp.Val {
				tmp = next
				break
			}
			tmp.Next = next.Next
		}
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
func _arrayToListNode(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	head := &ListNode{Val:a[0]}
	prev := head
	for i, v := range a{
		if i==0 {
			continue
		}
		prev.Next = &ListNode{Val:v}
		prev = prev.Next
	}
	return head
}
