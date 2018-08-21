package main

func main() {
	t := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	println("1->2->3->4,")
	printListNode(swapPairs(t))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pNow := head
	newHead := head.Next
	pPre := head
	for {
		if pNow == nil || pNow.Next == nil {
			break
		}
		pPre.Next = pNow.Next
		pNow.Next = pNow.Next.Next
		pPre.Next.Next = pNow
		pNow = pNow.Next

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
