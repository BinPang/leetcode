package main

func main() {
	println(countNumbersWithUniqueDigits(0))
	println(countNumbersWithUniqueDigits(1))
	println(countNumbersWithUniqueDigits(2))
	println(countNumbersWithUniqueDigits(3))
	println(countNumbersWithUniqueDigits(20))
}

func countNumbersWithUniqueDigits(n int) int {
	/*
		f(0) = 1
		f(1) = 9 + f(0) = 10
		f(2) = 9 * 9 + f(1) = 91
		f(3) = 9 * 9 * 8 + f(2) = 648 + 91 = 739

		0	1
		1	10
		2	91
		3	739
	*/
	arr := []int{9, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	if n == 0 {
		return 1
	}
	if n > 10 {
		n = 10
	}
	prev := 1
	fun := 1
	now := 0
	for i := 0; i < n; i++ {
		prev = prev * arr[i]
		now = prev + fun
		fun = now
	}
	return now
}
