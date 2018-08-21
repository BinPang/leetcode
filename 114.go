package main

func main() {
	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	flatten(t)
	print_flatten(t)

	t = &TreeNode{
		Val:1,
		Left:&TreeNode{
			Val:2,
			Left:&TreeNode{
				Val:3,
			},
		},
	}
	flatten(t)
	print_flatten(t)
}

func print_flatten(root *TreeNode) {
	println("start....")
	p := root
	for {
		if p == nil {
			break
		}
		print(p.Val, ",")
		if p.Left != nil {
			print("_, left is not nil,__")
		}
		p = p.Right
	}
	println("\nend...")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	_flatten(root)
}

func _flatten(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}
	var l *TreeNode
	var r *TreeNode
	if root.Left != nil {
		l = _flatten(root.Left)
	}
	if root.Right != nil {
		r = _flatten(root.Right)
	}
	if l != nil {
		l.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	if r != nil {
		return r
	} else {
		return l
	}
}
