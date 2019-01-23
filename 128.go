package main

func main() {
	println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func longestConsecutive(nums []int) int {
	data := map[int]bool{}
	for _, v := range nums {
		data[v] = true
	}
	r := 0
	sliceR := 0
	i := 0
	for _, v := range nums {
		if !data[v] {
			continue
		}
		data[v] = false
		sliceR = 1
		i = 1
		for {
			if data[v+i] {
				sliceR++
				data[v+i] = false
				i++
			} else {
				break
			}
		}
		i = 1
		for {
			if data[v-i] {
				sliceR++
				data[v-i] = false
				i++
			} else {
				break
			}
		}
		if sliceR > r {
			r = sliceR
		}
	}
	return r
}
