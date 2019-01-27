package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	removeElements(&ListNode{Val:1, Next:&ListNode{Val:2, Next:&ListNode{Val:3}}}, 3)
}

func removeElements(head *ListNode, val int) *ListNode {
	tmp := head
	for {
		if tmp == nil {
			return nil
		}
		if tmp.Val == val {
			tmp = tmp.Next
		} else {
			break
		}
	}
	r := tmp
	for {
		if tmp.Next == nil {
			break
		}
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}

	return r
}
