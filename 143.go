package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	reorderList(list)
	for list != nil {
		println(list.Val)
		list = list.Next
	}

	list = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next:&ListNode{Val:5}}}}}
	reorderList(list)
	for list != nil {
		println(list.Val)
		list = list.Next
	}
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	l := 0
	tmp1 := head
	tmp2 := head
	startOne := head
	startTwo := head
	end := head
	for {
		l += 1
		if tmp1.Next == nil {
			end = tmp1
			break
		}
		tmp1 = tmp1.Next
	}
	step := 0
	if l%2 == 0 {
		step = l / 2
	} else {
		step = l/2 + 1
	}
	for {
		if step == 1 {
			tmp1 = startTwo.Next
			startTwo.Next = nil
			startTwo = tmp1
			break
		}
		startTwo = startTwo.Next
		step--
	}
	tmp1 = startTwo //old start
	for startTwo != end {
		tmp2 = startTwo.Next
		startTwo.Next = end.Next
		end.Next = startTwo
		startTwo = tmp2
	}
	startTwo = end
	end = tmp1
	end.Next = nil

	for startTwo != nil {
		tmp1 = startOne.Next
		tmp2 = startTwo.Next
		startOne.Next = startTwo
		startTwo.Next = tmp1

		startOne = tmp1
		startTwo = tmp2
	}
}
