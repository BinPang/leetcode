package main

import "fmt"

func main() {
	println(fmt.Sprintf("%+v", maxSlidingWindow([]int{}, 0)))
	println(fmt.Sprintf("%+v", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)))
	println(fmt.Sprintf("%+v", maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 1)))
}

type deque struct {
	s int
	d []int
}

func (d *deque) Push(v int) {
	d.d = append(d.d, v)
	for i := len(d.d) - 2; i >= d.s; i-- {
		if d.d[i] < d.d[i+1] {
			d.d[i] = d.d[i+1]
			continue
		}
		break
	}
}
func (d *deque) Shift() int {
	d.s++
	return d.d[d.s-1]
}

func maxSlidingWindow(nums []int, k int) []int {
	l := len(nums)
	if l == 0 {
		return nil
	}
	r := make([]int, 0)
	d := deque{d: make([]int, 0)}
	max := 0
	for i := 0; i < k; i++ {
		d.Push(nums[i])
	}
	for i := k; i < l; i++ {
		max = d.Shift()
		r = append(r, max)
		d.Push(nums[i])
	}
	r = append(r, d.Shift())
	return r
}
