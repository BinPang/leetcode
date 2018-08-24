package main

func main() {
	println(removeElement([]int{3, 2, 2, 3}, 3))
	println(removeElement([]int{}, 3))
	println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	println(removeElement([]int{2}, 2))
	println(removeElement([]int{2, 2, 2, 2}, 2))
}

func removeElement(nums []int, val int) int {
	l := 0
	index := 0
	for _, v := range nums {
		if v != val {
			nums[index] = v
			index ++
			l ++
		}
	}
	return l
}
