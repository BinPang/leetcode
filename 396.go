package main

func main() {
	println(maxRotateFunction([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	println(maxRotateFunction([]int{4, 3, 2, 6}))
}

func maxRotateFunction(A []int) int {
	total := 0
	sum := 0
	l := len(A)
	r := 0
	for k, v := range A {
		total += v
		sum += k * v
	}
	r = sum
	for i := l-1; i > 0; i-- {
		sum = sum - l*A[i] + total
		if r < sum {
			r = sum
		}
	}
	return r
}
