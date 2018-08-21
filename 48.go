package main

func main() {
	t := [][]int{
		{1, 2, 3, 4, 5},
		{12, 13, 14, 15, 16},
		{23, 24, 25, 26, 27},
		{33, 34, 35, 36, 37},
		{43, 44, 45, 46, 47},
	}
	t1 := [][]int{
		{1, 2,3,4},
		{11, 12, 13, 14},
		{21, 22, 23, 24},
		{31, 32, 33, 34},
	}
	printMatrix(t)
	rotate48(t)
	printMatrix(t)

	printMatrix(t1)
	rotate48(t1)
	printMatrix(t1)
}

func printMatrix(matrix [][]int) {
	for _, v := range matrix {
		for _, v1 := range v {
			print(v1, ",")
		}
		println("")
	}
}

func rotate48(matrix [][]int) {
	l := len(matrix)
	for level := 0; level*2 < l-1; level++ {

		co := l - level*2
		for i := 0; i < co-1; i++ {

			now := matrix[level][level+i]
			matrix[level][level+i] = matrix[l-level-1-i][level]         //top-left
			matrix[l-level-1-i][level] = matrix[l-level-1][l-level-1-i] //bottom-left
			matrix[l-level-1][l-level-1-i] = matrix[level+i][l-level-1] // bottom-right
			matrix[level+i][l-level-1] = now                            //top-right
			//if i == 1 {return}
		}
		//return
	}
}
