package main

func main() {
	println("1->2->3->4,")
	printListNode(swapPairs(_arrayToListNode([]int{1, 2, 3, 4})))
	println("1->2->3->4->5,")
	printListNode(swapPairs(_arrayToListNode([]int{1, 2, 3, 4, 5})))
	println(",")
	printListNode(swapPairs(nil))
	println("1,")
	printListNode(swapPairs(_arrayToListNode([]int{1})))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = newHead.Next
	newHead.Next = head
	pPre := head
	pNow := head.Next
	for {
		if pNow == nil || pNow.Next == nil {
			break
		}
		pPre.Next = pNow.Next
		pNow.Next = pNow.Next.Next
		pPre.Next.Next = pNow
		pPre = pNow
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

