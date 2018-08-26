package main

func main() {
	println("input:[2,5, 4, 5, 1], need return:12,real:", largestRectangleArea([]int{2, 5, 4, 5, 1}))

	println("input:[2,1,5,6,2,2,2,2], need return:12,real:", largestRectangleArea([]int{2, 1, 5, 6, 2, 2, 2, 2}))
	println("input:[2,1,5,6,2,3], need return:10,real:", largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	println("input:[2,3,1,1,1,1,1,1,1], need return:9,real:", largestRectangleArea([]int{2, 3, 1, 1, 1, 1, 1, 1, 1}))
	println("input:[2,5, 5, 5, 1,5,6,2,3], need return:15,real:", largestRectangleArea([]int{2, 5, 5, 5, 1, 5, 6, 2, 3}))
	println("input:[1], need return:1,real:", largestRectangleArea([]int{1}))
	println("input:[1,1], need return:2,real:", largestRectangleArea([]int{1, 1}))
	println("input:[2,3,4,5,1], need return:9,real:", largestRectangleArea([]int{2, 3, 4, 5, 1}))
	println("input:[1,5,6], need return:10,real:", largestRectangleArea([]int{1, 5, 6}))

}

func largestRectangleArea(heights []int) int {
	l := len(heights)
	if l == 0 {
		return 0
	}
	stack := make([]int, l)
	index := 0
	r := 0
	min := 0
	step := 1
	loopMin := 0
	for i, v := range heights {
		if i == 0 {
			stack[i] = v
			r = v
			min = v
			continue
		}
		if min < v {
			min = v
			stack[i] = min
		}
		if v >= stack[i-1] {
			stack[i] = v
		} else {
			stack[i] = v
			step = 1
			index = i - 1
			//println("*", i, v, r)
			loopMin = stack[index]
			for {
				if index < 0 || v > stack[index] {
					break
				}
				if stack[index] < loopMin {
					loopMin = stack[index]
				}
				maxNow := loopMin * step
				//println("---", loopMin, maxNow)
				if maxNow > r {
					r = maxNow
				}
				step ++
				index --
			}
		}
	}

	//_printArray(stack)
	//println(r)
	index = l - 1
	step = 1
	for index = l - 1; index >= 0; index-- {
		if index != l-1 && stack[index] > stack[index+1] {
			stack[index] = stack[index+1]
		}
		maxNow := stack[index] * step
		if maxNow > r {
			r = maxNow
		}
		step++
	}
	return r
}

func _printArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}
