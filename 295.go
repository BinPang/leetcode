package main

import (
	"container/heap"
	"fmt"
)

func t1() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	obj.AddNum(3)
	println(int(obj.FindMedian()))
	println(fmt.Sprintf("%+v,%+v", obj.small, obj.big))
}
func t2() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	obj.AddNum(-1)
	println(int(obj.FindMedian()))
	println(fmt.Sprintf("%+v,%+v", obj.small, obj.big))
}
func t3() {
	obj := Constructor()
	obj.AddNum(40)
	obj.AddNum(12)
	obj.AddNum(16)
	println(int(obj.FindMedian()))
	println(fmt.Sprintf("%+v,%+v", obj.small, obj.big))
}
func t4() {
	obj := Constructor()
	obj.AddNum(-1)
	obj.AddNum(-2)
	obj.AddNum(-3)
	obj.AddNum(-4)
	obj.AddNum(-5)
	println(int(obj.FindMedian()))
	println(fmt.Sprintf("%+v,%+v", obj.small, obj.big))
}

func main() {
	t1()
	t2()
	t3()
	t4()
	return
}

type PriorityQueue struct {
	max  bool
	item []int
}

func (pq PriorityQueue) Len() int { return len(pq.item) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq.max {
		return pq.item[i] > pq.item[j]
	}
	return pq.item[i] < pq.item[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.item[i], pq.item[j] = pq.item[j], pq.item[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.item = append(pq.item, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	item := (pq.item)[len(pq.item)-1]
	pq.item = (pq.item)[0 : len(pq.item)-1]
	return item
}

type MedianFinder struct {
	small PriorityQueue
	big   PriorityQueue
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	pq1 := PriorityQueue{max: true, item: make([]int, 0)}
	heap.Init(&pq1)
	pq2 := PriorityQueue{max: false, item: make([]int, 0)}
	heap.Init(&pq2)
	return MedianFinder{small: pq1, big: pq2}
}

func (this *MedianFinder) AddNum(num int) {
	//println(fmt.Sprintf("%+v,%+v", this.small, this.big))
	if this.big.Len() > 0 {
		if this.small.Len() == this.big.Len() {
			if this.big.item[0] < num {
				heap.Push(&this.small, heap.Pop(&this.big))
				heap.Push(&this.big, num)
			} else {
				heap.Push(&this.small, num)
			}
		} else {
			if this.small.item[0] > num {
				heap.Push(&this.big, heap.Pop(&this.small))
				heap.Push(&this.small, num)
			} else {
				heap.Push(&this.big, num)
			}
		}
	} else {
		if this.small.Len() == 0 {
			heap.Push(&this.small, num)
		} else {
			old := heap.Pop(&this.small).(int)
			if old <= num {
				heap.Push(&this.small, old)
				heap.Push(&this.big, num)
			} else {
				heap.Push(&this.small, num)
				heap.Push(&this.big, old)
			}
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.small.Len() == 0 {
		return 0
	} else if this.big.Len() == 0 {
		return float64(this.small.item[0])
	}
	if this.small.Len() == this.big.Len() {
		return float64(this.small.item[0]+this.big.item[0]) / 2
	} else {
		return float64(this.small.item[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
