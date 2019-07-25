package main

func fib(N int) int {
	if N == 0 || N == 1 {
		return N
	}
	a, b := 0, 1
	for N > 1 {
		b = a + b
		a = b - a
		N--
	}
	return b
}
