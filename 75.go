package main

func main() {
	t3 := []int{1, 2, 0}
	println("input:[2,0,1],output:")
	sortColors(t3)
	_printArray(t3)
	return

	t0 := []int{2,0,1}
	println("input:[2,0,1],output:")
	sortColors(t0)
	_printArray(t0)


	t2 := []int{1, 0}
	println("input:[1, 0],output:")
	sortColors(t2)
	_printArray(t2)



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
	startZero := 0
	endTwo := l - 1
	for i, v := range nums {
		println(i, v, startZero, endTwo)
		if v == 0 {
			nums[i] = nums[startZero]
			nums[startZero] = v
			startZero ++
		} else if v == 2 {
			if i >=endTwo {
				break
			}
			nums[i] = nums[endTwo]
			nums[endTwo] = v
			endTwo --
		}

	}
}

func _printArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}
