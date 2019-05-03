package main

import "container/list"

type MyQueue struct {
	l *list.List
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	r := MyQueue{}
	r.l = list.New()
	return r
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.l.PushBack(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	tmp := this.l.Front()
	this.l.Remove(tmp)

	return tmp.Value.(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.l.Front().Value.(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.l.Len() == 0
}
