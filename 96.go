package main

func main() {
	println("n=3, result:", numTrees(3))
	println("n=4, result:", numTrees(4))
}

func numTrees(n int) int {
	m := map[int]int{
		0: 1,
		1: 1,
		2: 2,
	}
	if n < 0 {
		panic("n should gte 0")
	}
	if v, ok := m[n]; ok {
		return v
	}
	tmp := 0
	for i := 3; i <= n; i++ {
		tmp = 0
		for j := 1; j <= i; j++ {
			tmp += m[j-1] * m[i-j]
		}
		m[i] = tmp
	}
	return m[n]
}
