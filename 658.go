package main

import (
	"fmt"
)

func main() {
	println(fmt.Sprintf("%+v", findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3)))
	println(fmt.Sprintf("%+v", findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1)))
	println(fmt.Sprintf("%+v", findClosestElements([]int{1, 2, 3, 4, 5}, 4, 6)))
	println(fmt.Sprintf("%+v", findClosestElements([]int{2, 4, 6, 8, 10}, 4, 5)))
	println(fmt.Sprintf("%+v", findClosestElements([]int{2, 4, 6, 8, 10}, 4, 7)))
	println(fmt.Sprintf("%+v", findClosestElements([]int{2, 4, 6, 8, 10}, 4, 9)))
	println(fmt.Sprintf("%+v", findClosestElements([]int{2, 4, 6, 8, 10}, 4, 6)))
}

func findClosestElements(arr []int, k int, x int) []int {
	//find closest x index
	l := len(arr)
	start, end, middle := 0, len(arr)-1, 0
	for {
		middle = (start + end) / 2
		if start == middle {
			if x-arr[start] <= arr[end]-x {
				middle = start
			} else {
				middle = end
			}
			break
		}
		if arr[middle] == x {
			break
		}
		if arr[middle] < x {
			start = middle
		} else {
			end = middle
		}
	}
	start = middle - 1
	end = middle + 1
	for i := 1; i < k; i++ {
		if start < 0 {
			end++
			continue
		} else if end >= l {
			start--
			continue
		}
		if x-arr[start] <= arr[end]-x {
			start--
		} else {
			end++
		}
	}
	return arr[start+1 : end]
}
