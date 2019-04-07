package main

import "container/list"

type MyStack struct {
	l *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{l: list.New()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.l.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.l.Remove(this.l.Back()).(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.l.Back().Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.l.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
