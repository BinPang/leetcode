package main

import (
	"fmt"
	"math/rand"
)

func main() {
	row := 8
	col := 23
	obj := Constructor(row, col)
	test := map[string]int{}
	for i := 0; i < row*col; i++ {
		tt := fmt.Sprintf("%+v", obj.Flip())
		if test[tt] > 0 {
			println("___", i, test[tt], tt)
		} else {
			test[tt] = i
			//println("+++", i, test[tt], tt)
		}
	}
}

type Solution struct {
	rows   int
	cols   int
	remain int
	m      map[int]int
}

func Constructor(n_rows int, n_cols int) Solution {
	return Solution{
		rows:   n_rows,
		cols:   n_cols,
		remain: n_rows * n_cols,
		m:      make(map[int]int),
	}
}

func (this *Solution) Flip() []int {
	r := rand.Intn(this.remain) //[0, n)
	this.remain--

	origin, ok := this.m[r]
	if !ok {
		origin = r
	}

	if _, ok = this.m[this.remain]; ok {
		this.m[r] = this.m[this.remain]
	} else {
		this.m[r] = this.remain
	}
	//println(origin, r, this.remain, fmt.Sprintf("%+v", this.m))

	return []int{origin / this.cols, origin % this.cols}
}

func (this *Solution) Reset() {
	this.remain = this.rows * this.cols
	this.m = make(map[int]int)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n_rows, n_cols);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
