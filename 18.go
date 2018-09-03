package main

import "sort"

func main() {
	t0 := fourSum([]int{0, 0, 0, 0}, 0)
	_printArrayArray(t0)
	return
	t1 := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	_printArrayArray(t1)
}

func fourSum(nums []int, target int) [][]int {
	l := len(nums)
	if l <= 3 {
		return nil
	}

	sort.Ints(nums)
	//_printArray(nums)
	r := make([][]int, 0)
	start := 0
	end := 0
	needSum := 0
	for i := 0; i < l-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < l-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			start = j + 1
			end = l - 1
			needSum = target - nums[i] - nums[j]
			for start < end {
				if nums[start]+nums[end] == needSum {
					//println(i, j, start, end, nums[i], nums[j], nums[start], nums[end])
					r = append(r, []int{nums[i], nums[j], nums[start], nums[end]})
					for {
						start ++
						if start >= end || nums[start-1] != nums[start] {
							break
						}
					}
					for {
						end--
						if start >= end || nums[end] != nums[end+1] {
							break
						}
					}
				} else if nums[start]+nums[end] < needSum {
					for {
						start ++
						if start >= end || nums[start-1] != nums[start] {
							break
						}
					}
				} else {
					for {
						end--
						if start >= end || nums[end] != nums[end+1] {
							break
						}
					}
				}
			}
		}
	}
	return r
}

func _printArrayArray(t [][]int) {
	for _, v := range t {
		_printArray(v)
	}
}
func _printArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}
