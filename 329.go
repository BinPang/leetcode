package main

/**
Input: nums =
[
  [9,9,4],
  [6,6,8],
  [2,1,1]
]
Output: 4
Explanation: The longest increasing path is [1, 2, 6, 9].
*/
func longestIncreasingPath(matrix [][]int) int {
	r := 0
	l := len(matrix)
	if l == 0 {
		return 0
	}
	l1 := len(matrix[0])
	if l1 == 0 {
		return 0
	}
	tmp := make([][]int, l)
	for i := 0; i < l; i++ {
		tmp[i] = make([]int, l1)
	}
	startX := 0
	startY := 0
	for  {

	}

	return r
}
