package main

func main() {
	println(fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
	println(fourSumCount([]int{1, 2, -1}, []int{-2, -1, 1}, []int{-1, 2, 0}, []int{0, 2, 1}))
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
	r := 0
	t1 := make(map[int]int, 2*len(A))
	for _, v := range A {
		for _, v1 := range B {
			t1[v+v1] += 1
		}
	}
	for _, v := range C {
		for _, v1 := range D {
			if t1[-v-v1] > 0 {
				r += t1[-v-v1]
			}
		}
	}

	return r
}
