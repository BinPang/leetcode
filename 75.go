package main

func main() {

	t2 := []int{1, 0}
	println("input:[1, 0],output:")
	sortColors(t2)
	_printArray(t2)

	t0 := []int{2, 0, 1}
	println("input:[2,0,1],output:")
	sortColors(t0)
	_printArray(t0)


	t3 := []int{1, 2, 0}
	println("input:[1, 2, 0],output:")
	sortColors(t3)
	_printArray(t3)

	t := []int{2, 0, 2, 1, 1, 0}
	println("input:[2,0,2,1,1,0],output:")
	sortColors(t)
	_printArray(t)

	t1 := []int{2, 0, 2, 1, 1, 0, 0}
	println("input:[2,0,2,1,1,0, 0],output:")
	sortColors(t1)
	_printArray(t1)
}

func sortColors(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}
	oneStart := -1
	start := 0
	end := l - 1
	for {
		if start == end {
			if nums[start] == 0 {
				if oneStart == -1 {
					oneStart = 0
				}
				nums[oneStart], nums[start] = nums[start], nums[oneStart]
			}
			break
		}
		//println("se:", start, ",", end)
		if nums[start] == 0 {
			if oneStart != -1 {
				nums[start], nums[oneStart] = nums[oneStart], nums[start]
				oneStart ++
			}
		} else if nums[start] == 1 {
			if oneStart == -1 {
				oneStart = start
			}
		} else {
			for {
				if nums[end] != 2 {
					break
				}
				end--
				if start == end {
					return
				}
			}
			nums[start], nums[end] = nums[end], nums[start]
			if nums[start] == 0 {
				if oneStart != -1 {
					nums[start], nums[oneStart] = nums[oneStart], nums[start]
					oneStart ++
				}
			} else {
				if oneStart == -1 {
					oneStart = start
				}
			}
		}
		start++
	}
}

func _printArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}
