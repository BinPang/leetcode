package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmpResult := result
	moreTen := 0
	for {
		tmpPlus := l1.Val + l2.Val + moreTen
		if tmpPlus >= 10 {
			tmpPlus -= 10
			moreTen = 1
		} else {
			moreTen = 0
		}
		tmpResult.Val = tmpPlus
		if l1.Next == nil && l2.Next == nil {
			if moreTen == 1 {
				tmpResult.Next = &ListNode{moreTen, nil}
			}
			break
		}
		if l1.Next == nil {
			l1.Val = 0
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}
		tmpResult.Next = &ListNode{}
		tmpResult = tmpResult.Next
	}
	return result
}
