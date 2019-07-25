package main

import (
	"container/heap"
	"fmt"
)

/*
(1,2) -> (1,9) -> (1,10) -> (1,15)
(7,2) -> (7,9) -> (7,10) -> (7,15)
(11,2) -> (11,9) -> (11,10) -> (11,15)
(16,2) -> (16,9) -> (16,10) -> (16,15)
*/

func main() {
	//my:[[1,1],[1,1],[1,2],[2,1],[2,2],[2,3]]
	//ex:[[1,1],[1,1],[2,1],[1,2],[1,2],[2,2],[1,3],[1,3],[2,3]]
	println(fmt.Sprintf("%+v", kSmallestPairs([]int{1, 7, 11, 16}, []int{2, 9, 10, 15}, 6)))
	println(fmt.Sprintf("%+v", kSmallestPairs([]int{1, 1, 1}, []int{1, 1, 1}, 20)))
	println(fmt.Sprintf("%+v", kSmallestPairs([]int{1, 3, 3}, []int{2, 4, 4}, 20)))
	println(fmt.Sprintf("%+v", kSmallestPairs([]int{1, 3, 3}, []int{2, 4, 4}, 5)))
	println(fmt.Sprintf("%+v", kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3)))
	println(fmt.Sprintf("%+v", kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2)))
	println(fmt.Sprintf("%+v", kSmallestPairs([]int{1, 2}, []int{3}, 3)))
}

type Item struct {
	one      int
	two      int
	priority int // The priority of the item in the queue.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	//return pq[i].priority > pq[j].priority
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return item
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 == 0 || l2 == 0 {
		return nil
	}
	// Init heap
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for i := 0; i < l1; i++ {
		heap.Push(&pq, &Item{i, 0, nums1[i] + nums2[0]})
	}
	if k > l1*l2 {
		k = l1 * l2
	}
	r := make([][]int, k)
	for i := 0; i < k; i++ {
		item := heap.Pop(&pq).(*Item)
		r[i] = []int{nums1[item.one], nums2[item.two]}
		if item.two < l2-1 {
			heap.Push(&pq, &Item{item.one, item.two + 1, nums1[item.one] + nums2[item.two+1]})
		}
	}
	return r
}
