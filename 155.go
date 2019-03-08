package main

func main() {
	minStack := Constructor()
	//minStack.Push(-2)
	//minStack.Push(0)
	//minStack.Push(-3)
	//println(minStack.GetMin())
	//minStack.Pop()
	//println(minStack.Top())
	//println(minStack.GetMin())

	minStack.Push(-10)
	minStack.Push(14)
	minStack.Push(-20)
	minStack.Pop()
	println(minStack.Top())
	println(minStack.GetMin())
}

type MinStack struct {
	Min   int
	Stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{Stack: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	if len(this.Stack) == 0 {
		this.Stack = append(this.Stack, x, x)
		this.Min = x
	} else {
		if this.Min > x {
			this.Stack = append(this.Stack, x, x)
			this.Min = x
		} else {
			this.Stack = append(this.Stack, x, this.Min)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.Stack) == 0 {
		return
	}
	this.Stack = this.Stack[:len(this.Stack)-2]
	if len(this.Stack) > 0 {
		this.Min = this.Stack[len(this.Stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.Stack) == 0 {
		return 0
	}
	return this.Stack[len(this.Stack)-2]
}

func (this *MinStack) GetMin() int {
	if len(this.Stack) == 0 {
		return 0
	}
	return this.Stack[len(this.Stack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor()
 * obj.Push(x)
 * obj.Pop()
 * param_3 := obj.Top()
 * param_4 := obj.GetMin()
 */
