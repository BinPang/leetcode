package main

func main() {
	println("input:[1,1],need:1,result:", maxArea([]int{1,1}))
	println("input:[1,2],need:1,result:", maxArea([]int{1,2}))
	println("input:[1,8,6,2,5,4,8,3,7],need:49,result:", maxArea([]int{1,8,6,2,5,4,8,3,7}))
}

func maxArea(height []int) int {
	maxNow := 0
	start := 0
	end := len(height)-1
	tmpArea := 0
	for {
		if end-start < 1 {
			break
		}
		if height[start] <= height[end] {
			tmpArea = height[start] * (end-start)
			start +=1
		} else {
			tmpArea = height[end] * (end-start)
			end-=1
		}
		if maxNow < tmpArea {
			maxNow = tmpArea
		}
	}
	return maxNow
}