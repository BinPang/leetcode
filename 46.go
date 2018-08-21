package main

func main() {
	t := permute([]int{1, 2, 3, 4})
	for _, v := range t{
		for _, v1 := range v{
			print(v1, ",")
		}
		println("")
	}
}

func permute(nums []int) [][]int {
	l := len(nums)
	if l == 0 {
		return  nil
	}
	if l == 1 {
		r := make([][]int, 1)
		r[0] = make([]int, 1)
		r[0][0] = nums[0]
		return r
	}
	if l == 2 {
		r := make([][]int, 2)
		r[0] = []int{nums[0], nums[1]}
		r[1] = []int{nums[1], nums[0]}
		return r
	}

	r := make([][]int, 0)
	for i, v := range nums{
		small := make([]int, l-1)
		for j := range nums {
			if i==j {
				continue
			} else if i>j{
				small[j] = nums[j]
			} else {
				small[j-1] = nums[j]
			}
		}
		s := permute(small)
		for _, sv := range s{
			r = append(r, append([]int{v}, sv...))
		}
	}
	return r
}
