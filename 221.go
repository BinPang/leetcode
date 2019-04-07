package main

func main() {
	println(maximalSquare([][]byte{{'0', '0', '0'}, {'0', '0', '0'}, {'0', '0', '0'}, {'0', '0', '0'}}))

	println(maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))
}

func maximalSquare(matrix [][]byte) int {
	l1 := len(matrix)
	if l1 == 0 {
		return 0
	}
	l2 := len(matrix[0])
	if l2 == 0 {
		return 0
	}
	r := 0
	tmpR := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		tmpR[i] = make([]int, l2+1)
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if matrix[i-1][j-1] == '1' {
				tmpR[i][j] = min(tmpR[i-1][j-1], tmpR[i][j-1], tmpR[i-1][j]) + 1
				if r < tmpR[i][j] {
					r = tmpR[i][j]
				}
			}
		}
	}
	return r * r
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		//b < a
		if b < c {
			return b
		} else {
			return c
		}
	}
}
