package main

func main() {
	t4 := []int{2,3,1}
	_printArray(t4)
	nextPermutation(t4)
	_printArray(t4)

	t3 := []int{1,2, 7,6,5,4,3}
	_printArray(t3)
	nextPermutation(t3)
	_printArray(t3)


	t0 := []int{1,2,3}
	_printArray(t0)
	nextPermutation(t0)
	_printArray(t0)

	t1 := []int{3,2,1}
	_printArray(t1)
	nextPermutation(t1)
	_printArray(t1)

	t2 := []int{1,1,5}
	_printArray(t2)
	nextPermutation(t2)
	_printArray(t2)
}

func nextPermutation(nums []int) {
	l := len(nums)
	down := true
	for i, v := range nums {
		if i == 0 {
			continue
		}
		if v > nums[i-1] {
			down = false
			break
		}
	}
	if down {
		start := 0
		end := l - 1
		for {
			if start >= end {
				return
			}
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}
	start := l - 1
	for i := l - 2; i >= 0; i-- {
		if nums[i] < nums[start] {
			nums[i], nums[start] = nums[start], nums[i]
			j := i+1
			for {
				if j>=start {
					break
				}
				nums[j], nums[start] = nums[start], nums[j]
				start--
				j++
			}
			break
		} else if nums[i] < nums[i+1] {
			j := i
			for {
				if nums[j] > nums[i] && nums[j+1]<=nums[i]{
					break
				}
				j+=1
			}
			nums[i], nums[j] = nums[j], nums[i]
			i +=1
			for {
				if i >= start {
					break
				}
				nums[i], nums[start] = nums[start], nums[i]
				start--
				i++
			}
			break
		}
	}

}

func _printArray(t []int)  {
	for _, v := range t{
		print(v, ",")
	}
	println("")
}
