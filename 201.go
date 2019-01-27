package main

func main() {
	println(rangeBitwiseAnd(7, 8))
	println(rangeBitwiseAnd(5, 7))
	println(rangeBitwiseAnd(0, 1))
	println(rangeBitwiseAnd(11, 13))
}

func rangeBitwiseAnd(m int, n int) int {
	if m == 0 {
		return 0
	}
	if n == 1 {
		return m & n
	}
	if m == n {
		return m
	}
	m0, n0 := m, n
	for {
		if m0 == 0 && n0 == 0 {
			break
		}
		if m0 == 0 || n0 == 0 {
			return 0
		}
		m0 >>= 1
		n0 >>= 1
	}
	m0 = m
	n0 = 1
	for m0 >= 2 {
		m0 >>= 1
		n0 <<= 1
	}
	return n0 + rangeBitwiseAnd(m-n0, n-n0)
}
