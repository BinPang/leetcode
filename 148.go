package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 Sort a linked list in O(n log n) time using constant space complexity.
 Input: -1->5->3->4->0
Output: -1->0->3->4->5
  */
func main() {
	//origin := _arrayToListNode([]int{-1, 5, 3, 4, 0})
	origin := _arrayToListNode([]int{-1, 5})
	printListNode(sortList(origin))

	origin1 := _arrayToListNode([]int{4, 19, 14, 5, -3, 1, 8, 5, 11, 15})
	printListNode(sortList(origin1))
	println("___________")
	return
}

func sortList(head *ListNode) *ListNode {
	//empty or one item return
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		prev = slow
		slow = slow.Next
	}
	prev.Next = nil

	head = sortList(head)
	slow = sortList(slow)

	return merge(head, slow)
}

func merge(one *ListNode, two *ListNode) *ListNode {
	head := &ListNode{}
	tmp := head
	for one != nil && two != nil {
		if one.Val < two.Val {
			tmp.Next = one
			one = one.Next
		} else {
			tmp.Next = two
			two = two.Next
		}
		tmp = tmp.Next
	}
	if one != nil {
		tmp.Next = one
	}
	if two != nil {
		tmp.Next = two
	}

	return head.Next
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
