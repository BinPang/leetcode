package main

func main() {
	println(thirdMax([]int{1, 1, 1}))
	println(thirdMax([]int{1, 1, 2, 2, 2, 2}))
	println(thirdMax([]int{2, 2, 3, 1}))
	println(thirdMax([]int{2, 2, 3, 1, 4}))
	println(thirdMax([]int{2, 2, 3, 1, 4, 5, 6, 7}))
}

func thirdMax(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	one, two, three := nums[0], 0, 0 //one is max
	index := 1
	//find two
	for {
		if index == l {
			return one //not find the second
		}
		if nums[index] == one {
			index++
			continue
		}
		if nums[index] > one {
			two = one
			one = nums[index]
		} else {
			two = nums[index]
		}
		break
	}
	//find three
	findThird := false
	for {
		if index == l {
			if findThird {
				return three
			} else {
				return one
			}
		}
		if nums[index] == one || nums[index] == two {
			index++
			continue
		}
		findThird = true
		if nums[index] > one {
			three = two
			two = one
			one = nums[index]
		} else if nums[index] > two {
			three = two
			two = nums[index]
		} else {
			three = nums[index]
		}
		break
	}
	for index < l {
		if nums[index] == one || nums[index] == two || nums[index] == three {
			index++
			continue
		}
		if nums[index] > one {
			three = two
			two = one
			one = nums[index]
		} else if nums[index] > two {
			three = two
			two = nums[index]
		} else if nums[index] > three {
			three = nums[index]
		}
		index++
	}
	return three
}
