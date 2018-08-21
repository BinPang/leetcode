package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
func main() {
	ne := removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2)
	print("list:1,2,3,4,5,return:")
	printListNode(ne)

	ne = removeNthFromEnd(nil, 0)
	print("list:,return:")
	printListNode(ne)

	ne = removeNthFromEnd(&ListNode{1, nil}, 1)
	print("list:1,return:")
	printListNode(ne)

	ne = removeNthFromEnd(&ListNode{1, &ListNode{2, nil}}, 2)
	print("list:1,2,return:")
	printListNode(ne)
}

type ListNode struct {
	Val  int
	Next *ListNode
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	nodeOne := head
	var nodeTwo *ListNode
	for {
		n-=1
		if n <= 0 {
			break
		}
		nodeOne = nodeOne.Next
	}
	for {
		if nodeOne.Next == nil {
			if nodeTwo == nil {
				head = head.Next
				break
			}
			nodeTwo.Next = nodeTwo.Next.Next
			break
		}
		if nodeTwo == nil {
			nodeTwo = head
		} else {
			nodeTwo = nodeTwo.Next
		}
		nodeOne = nodeOne.Next
	}
	return head
}

//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	if head == nil {
//		return head
//	}
//	type arList struct {
//		before *ListNode
//		value  int
//		after  *ListNode
//	}
//	start := head
//	var prev *ListNode
//	ar := make([]arList, 0)
//	for {
//		if start == nil {
//			break
//		}
//		ar = append(ar, arList{before: prev, value: start.Val, after: start.Next})
//		prev = start
//		start = start.Next
//	}
//	if ar[len(ar)-n].before != nil {
//		ar[len(ar)-n].before.Next = ar[len(ar)-n].after
//	} else{
//		head = head.Next
//	}
//
//	return head
//}
