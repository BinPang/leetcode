package main

import "fmt"

func main() {
	i1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	println(fmt.Sprintf("%+v", findDiagonalOrder(i1)))

	i2 := [][]int{
		{1, 2, 3, 4,5,6,7},
	}
	println(fmt.Sprintf("%+v", findDiagonalOrder(i2)))

	i3 := [][]int{
		{1},
		{2},
		{3},
		{4},
		{5},
	}
	println(fmt.Sprintf("%+v", findDiagonalOrder(i3)))

	i4 := [][]int{
		{1, 2, 3,4,5,6},
		{7, 8, 9,10,11,12},
	}
	println(fmt.Sprintf("%+v", findDiagonalOrder(i4)))
}

func findDiagonalOrder(matrix [][]int) []int {
	l := len(matrix)
	if l == 0 {
		return nil
	}
	l1 := len(matrix[0])
	if l1 == 0 {
		return nil
	}
	rmax := l - 1
	cmax := l1 - 1
	total := l * l1
	r := make([]int, total)
	startX := 0
	startY := 0
	toTop := true
	index := 0
	for index < total {
		r[index] = matrix[startX][startY]
		if toTop {
			if startX == 0 { //reach top
				toTop = false
				if startY == cmax { //reach right
					startX++
				} else {
					startY++
				}
			} else if startY == cmax { //reach right
				toTop = false
				startX++
			} else {
				startX--
				startY++
			}
		} else {
			if startX == rmax { //reach bottom
				toTop = true
				startY++
			} else if startY == 0 { //reach left
				toTop = true
				startX++
			} else {
				startX++
				startY--
			}
		}
		index++
	}
	return r
}
