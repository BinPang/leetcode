package main

import "strconv"

func main() {
	println(getPermutation(8, 678))
	return
	println(getPermutation(4, 10))
	println(getPermutation(4, 9))
	println(getPermutation(4, 24))
}

func getPermutation(n int, k int) string {
	//Given n will be between 1 and 9 inclusive.
	r := 0
	tmp := make([]int, n+1)
	got := map[int]bool{}
	tmp[0] = 0
	tmp[1] = 1
	for i := 2; i <= n; i++ {
		tmp[i] = i * tmp[i-1]
		println("start:", i, tmp[i])
	}
	t0 := 0
	t1 := 0
	for i := n; i > 1; i-- {
		t0 = k / tmp[i-1] //商
		t1 = k % tmp[i-1] //余数
		//println(t0, t1)
		if t1 == 0 {
			t0 -= 1
			t1 = tmp[i-1]
		}
		//println(t0, t1)
		t0+=1
		for j := 1; j <= t0; j++ {
			if got[j] {
				t0 += 1
			}
		}
		got[t0-1] = true
		t0++
		r = r*10 + t0
		k = t1
		println("___", r, k, tmp[i-1], t0, t1)
	}
	for i := 0; i < n; i++ {
		if !got[i] {
			r = r*10 + (i + 1)
			break
		}
	}
	return strconv.Itoa(r)
}
