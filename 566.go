package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m := len(nums)
	n := len(nums[0])
	if m*n != r*c {
		return nums
	}
	re := make([][]int, r)
	for i := 0; i < r; i++ {
		re[i] = make([]int, c)
	}
	i1 := 0
	i2 := 0
	for _, v := range nums {
		for _, vv := range v {
			re[i1][i2] = vv
			i2++
			if i2 == c {
				i1++
				i2 = 0
			}
		}
	}

	return re
}
