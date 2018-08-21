package main

func main() {
	println("[3,2,1,0,4],result:", canJump([]int{3, 2, 1, 0, 4}))
	println("[2, 0, 0, 0],result:", canJump([]int{2, 0, 0, 0}))
	println("[2, 0, 0],result:", canJump([]int{2, 0, 0}))
	println("[2, 0],result:", canJump([]int{2, 0}))
	println("[2,3,1,1,4],result:", canJump([]int{2, 3, 1, 1, 4}))
}

func canJump(nums []int) bool {
	l := len(nums)
	if l == 0 {
		return false
	}
	if l == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	reachMax := 0
	for i, v := range nums {
		if i == 0 {
			reachMax = v
			continue
		}
		maxNow := 0
		if v != 0 {
			maxNow = i + v
		}
		//println(i, v, maxNow, reachMax)
		if maxNow > reachMax {
			reachMax = maxNow
		}
		if v == 0 && reachMax == i && i != l-1 {
			return false
		}
	}
	return true
}
