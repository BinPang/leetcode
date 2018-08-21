package main

func main() {

	t2 := mergeKLists([]*ListNode{
		{Val:1},
	})
	printListNode(t2)
	return
	t := mergeKLists([]*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		{Val: 2, Next: &ListNode{Val: 6}},
	})
	printListNode(t)
	t1 := mergeKLists([]*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		{Val: 2, Next: &ListNode{Val: 6}},
		{Val: 3},
	})
	printListNode(t1)

	mergeKLists(nil)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	//even three for loop,we only use mn*lg2(mn),if m==n,then n^2*lg2(n^2)=n^2*lg2(n)
	l := len(lists)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return lists[0]
	}
	max := l - 1
	for {
		i := 0
		for {
			if lists[i] == nil {
				lists[i] = lists[max]
			} else if lists[max] == nil {
				//item 2 is nil, do nothing
			} else {
				if lists[i].Val > lists[max].Val {
					lists[i], lists[max] = lists[max], lists[i]
				}
				head1 := lists[i]
				head2 := lists[max]
				prev := lists[i]
				for {
					if head1 == nil {
						prev.Next = head2
						break
					} else if head2 == nil {
						break
					}
					if head1.Val <= head2.Val {
						prev = head1
						head1 = head1.Next
					} else {
						prev.Next = head2
						prev = head2
						head2 = head2.Next
						prev.Next = head1
					}
				}
			}
			i++
			max--
			if i >= max {
				break
			}
		}
		if max < 1 {
			break
		}
	}
	return lists[0]
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
