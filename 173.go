package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	data []*TreeNode
	len  int
}

func Constructor(root *TreeNode) BSTIterator {
	r := BSTIterator{
		data: make([]*TreeNode, 0),
		len:  0,
	}
	r.push(root)
	return r
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	tmp := this.data[this.len-1]
	this.len--
	this.push(tmp.Right)

	return tmp.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.len != 0
}

func (this *BSTIterator) push(node *TreeNode) {
	tmp := node
	for tmp != nil {
		if len(this.data) == this.len {
			this.data = append(this.data, tmp)
			this.len++
		} else {
			this.data[this.len] = tmp
			this.len++
		}
		tmp = tmp.Left
	}
}
