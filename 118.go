package main

func main() {
	generate(5)
}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	r := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		tmp := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				tmp[j] = 1
			} else {
				tmp[j] = r[i-1][j-1] + r[i-1][j]
			}
		}
		r[i] = tmp
	}
	return r
}
