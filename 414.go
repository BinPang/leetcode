package main

func main() {
	println(thirdMax([]int{1, 1, 1}))
	println(thirdMax([]int{1, 1, 2, 2, 2, 2}))
}

func thirdMax(nums []int) int {
	l := len(nums)
	tmp := make([]int, 3)
	l := 0
	for _, v := range nums {
		if l == 0 {
			tmp[0] = v
			l++
			continue
		}
		if tmp[0] == v {
			continue
		}
		if l == 1 {
			tmp[1] = v
			l++
			continue
		}
	}
}