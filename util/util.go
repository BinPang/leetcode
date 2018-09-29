package main

func main() {
	inOrder(_arrayToTreeNodeLeetCode([]interface{}{3, 1, nil, nil, 2}))
	inOrder(_arrayToTreeNodeLeetCode([]interface{}{3, 1, nil, nil, 2}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

//compatible with leetcode
func _arrayToTreeNodeLeetCode(a []interface{}) *TreeNode {
	l := len(a)
	if l == 0 || a[0] == nil {
		return nil
	}
	r := &TreeNode{Val:a[0].(int)}
	lastLevel := make([]*TreeNode, 0)//last level *TreeNode
	lastLevel = append(lastLevel, r)
	lastIndex := 0
	nowLevel := make([]*TreeNode, 0)
	i := 1
	for {
		if i >= l {
			break
		}
		if len(lastLevel) == 0 {
			break
		}
		if a[i] == nil {
			if lastIndex%2 == 0 {
				lastLevel[lastIndex/2].Left = nil
			} else {
				lastLevel[lastIndex/2].Right = nil
			}
		} else {
			tmp := &TreeNode{Val:a[i].(int)}
			nowLevel = append(nowLevel, tmp)
			if lastIndex%2 == 0 {
				lastLevel[lastIndex/2].Left = tmp
			} else {
				lastLevel[lastIndex/2].Right = tmp
			}
		}
		lastIndex ++
		i++
		if lastIndex == len(lastLevel)*2 {
			lastIndex = 0
			lastLevel = nowLevel
			nowLevel = make([]*TreeNode, 0)
		}
	}
	return r
}

func _arrayToTreeNode(a []interface{}) *TreeNode {
	t := make([]*TreeNode, len(a))
	for i, v := range a {
		t[i] = &TreeNode{}
		if i == 0 {
			if v == nil {
				return nil
			}
		} else {
			before := (i+1)/2 - 1
			lr := (i + 1) % 2
			if t[before] == nil {
				println("before is nil,i=", i, ",before=", before)
				panic("panic")
			}
			if v == nil {
				t[i] = nil
			}
			if lr == 0 {
				t[before].Left = t[i]
			} else {
				t[before].Right = t[i]
			}
		}
		if v != nil {
			t[i].Val = v.(int)
		}
	}
	return t[0]
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

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	print(node.Val, ",")
	inOrder(node.Left)
	inOrder(node.Right)
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

func PrintArrayArray(t [][]int) {
	for _, v := range t {
		PrintArray(v)
	}
}
func PrintArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
