package main

type ListNode struct {
	Val  int
	Next *ListNode
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

func main() {
	p2 := _arrayToListNode([]int{8, 9, 9})
	p22 := _arrayToListNode([]int{2})
	t2 := addTwoNumbers(p2, p22)
	println(t2)

	p0 := _arrayToListNode([]int{7, 2, 4, 3})
	p00 := _arrayToListNode([]int{5, 6, 4})
	t0 := addTwoNumbers(p0, p00)
	println(t0)

	p1 := _arrayToListNode([]int{7, 2, 4, 3})
	p11 := _arrayToListNode([]int{3, 5, 6, 4})
	t1 := addTwoNumbers(p1, p11)
	println(t1)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	len1 := 0
	tmp1 := l1
	for tmp1 != nil {
		len1++
		tmp1 = tmp1.Next
	}
	len2 := 0
	tmp2 := l2
	for tmp2 != nil {
		len2++
		tmp2 = tmp2.Next
	}
	if len1 < len2 {
		l1, l2 = l2, l1
		len1, len2 = len2, len1
	}
	arr1 := make([]*ListNode, len1-len2)
	arr2 := make([]*ListNode, len2*2)
	tmp1 = l1
	tmp2 = l2
	index := 0
	for len1 > len2+index {
		arr1[index] = tmp1
		tmp1 = tmp1.Next
		index++
	}
	index = 0
	for tmp1 != nil {
		arr2[index] = tmp1
		arr2[index+1] = tmp2
		tmp1 = tmp1.Next
		tmp2 = tmp2.Next
		index += 2
	}
	plus := 0
	index = len2*2 - 1
	for index > 0 {
		arr2[index-1].Val += arr2[index].Val + plus
		if arr2[index-1].Val >= 10 {
			arr2[index-1].Val -= 10
			plus = 1
		} else {
			plus = 0
		}
		index -= 2
	}
	index = len1 - len2 - 1
	for index >= 0 {
		arr1[index].Val = arr1[index].Val + plus
		if arr1[index].Val == 10 {
			arr1[index].Val -= 10
			plus = 1
		} else {
			plus = 0
			break
		}
		index--
	}
	if plus == 0 {
		return l1
	} else {
		return &ListNode{Val: 1, Next: l1}
	}
}
