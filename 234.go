package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	t0 := &ListNode{Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		}}
	println(isPalindrome(t0))
}

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	if fast == nil || fast.Next == nil {
		return true
	}
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			reverse(slow)
			slow = slow.Next
			break
		} else {
			slow = slow.Next
			fast = fast.Next.Next
		}
	}
	for slow != nil {
		if head.Val != slow.Val {
			return false
		}
		slow = slow.Next
		head = head.Next
	}
	return true
}

func reverse(head *ListNode) {
	item := head.Next
	if item == nil {
		return
	}
	for {
		if item.Next == nil {
			break
		}
		tmp := item.Next
		item.Next = item.Next.Next
		tmp.Next = head.Next
		head.Next = tmp
	}
}
