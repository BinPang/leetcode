package main

import (
	"container/heap"
	"fmt"
)

func main() {
	println(fmt.Sprintf("%+v", topKFrequent([]int{1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4}, 3)))
	println(fmt.Sprintf("%+v", topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)))
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue [][2]int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][1] < pq[j][1]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.([2]int))
}

func (pq *PriorityQueue) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return item
}
func topKFrequent(nums []int, k int) []int {
	tmp := make(map[int]int, 0)
	for _, v := range nums {
		tmp[v] += 1
	}
	pq := make(PriorityQueue, k)
	index := 0
	for kk, vv := range tmp {
		if index < k {
			pq[index] = [2]int{kk, vv}
		} else {
			if index == k {
				heap.Init(&pq)
			}
			if vv > pq[0][1] {
				pq[0] = [2]int{kk, vv}
				heap.Fix(&pq, 0)
			}
		}
		index++
	}
	r := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		v := heap.Pop(&pq).([2]int)
		r[i] = v[0]
	}
	return r
}
