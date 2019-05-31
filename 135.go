package main

func main() {
	println("14=", candy([]int{1, 0, 1, 0, 1, 1, 0, 1, 1}))
	println(candy([]int{1, 6, 5, 4, 0, 2}))
	println("12=", candy([]int{6, 5, 4, 0, 2}))
	println("5=", candy([]int{1, 0, 2}))
	println("4=", candy([]int{1, 2, 2}))
}

func candy(ratings []int) int {
	l := len(ratings)
	if l == 0 {
		return 0
	} else if l == 1 {
		return 1
	}
	r := 1
	prev := 1
	down := 0

	for i := 1; i < l; i++ {
		if ratings[i] > ratings[i-1] {
			if down != 0 {
				prev = 1
			}
			prev += 1
			r += prev
			down = 0
		} else if ratings[i] == ratings[i-1] {
			prev = 1
			r += prev
			down = 0
		} else {
			down += 1
			if down == prev {
				down += 1
			}
			r += down
		}
	}

	return r
}
