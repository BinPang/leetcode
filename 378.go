package main

import "container/heap"

func main() {
	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	println(kthSmallest(matrix, 8))
	println(kthSmallest(matrix, 1))
	println(kthSmallest(matrix, 4))
	println(kthSmallest(matrix, 6))
}

// A PriorityQueue implements heap.Interface and holds Items.
type PQ [][]int

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	//return pq[i].priority > pq[j].priority
	return pq[i][2] < pq[j][2]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PQ) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return item
}

func kthSmallest(matrix [][]int, k int) int {
	l := len(matrix)
	pq := PQ{}
	for i := 0; i < l; i++ {

		pq = append(pq, []int{i, 0, matrix[i][0]})
	}
	heap.Init(&pq)

	for k > 0 {
		v := heap.Pop(&pq).([]int)
		k--
		if k == 0 {
			return v[2]
		}
		v[1]++
		if v[1] < l {
			v[2] = matrix[v[0]][v[1]]
			heap.Push(&pq, v)
		}
	}
	return 0
}
