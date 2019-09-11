package main

func main() {
	println(findCircleNum([][]int{{1}}))
	println(findCircleNum([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}))

	println(findCircleNum([][]int{
		{1, 0, 0, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 0, 1, 1},
	}))

	println(findCircleNum([][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}))
}

func findCircleNum(M [][]int) int {
	r := 0
	N := len(M)
	for i := 0; i < N; i++ {
		if M[i][i] == 1 {
			r++
			M[i][i] = 0
			tmp := make([]int, 0)
			for j := 0; j < N; j++ {
				if M[i][j] == 1 && M[j][j] == 1 {
					M[j][j] = 0
					tmp = append(tmp, j)
				}
			}
			for len(tmp) > 0 {
				it := tmp[0]
				M[it][it] = 0
				for j := 0; j < N; j++ {
					if M[it][j] == 1 && M[j][j] == 1 {
						M[j][j] = 0
						tmp = append(tmp, j)
					}
				}
				tmp = tmp[1:]
			}
		}
	}
	return r
}
