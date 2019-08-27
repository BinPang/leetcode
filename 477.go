package main

func main() {
	println(totalHammingDistance([]int{4, 14, 2}))
	println(totalHammingDistance([]int{0, 0, 0}))
	println(totalHammingDistance([]int{0, 1, 1, 1}))
}

func totalHammingDistance(nums []int) int {
	r := 0
	l := len(nums)
	bitCount := 0
	var j uint
	for j = 0; j < 32; j++ {
		bitCount = 0
		for i := 0; i < l; i++ {
			bitCount += (nums[i] >> j) & 1
		}
		r += bitCount * (l - bitCount)
	}
	return r
	/*r := 0
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	l := 0
	for max > 0 {
		l++
		max >>= 1
	}
	number0 := make([]int, l)
	number1 := make([]int, l)
	tmp := 0
	index := 0

	for _, v := range nums {
		tmp = v
		index = 0
		for tmp > 0 {
			if tmp&1 == 1 {
				number1[index] += 1
			} else {
				number0[index] += 1
			}
			tmp >>= 1
			index++
		}
		for index < l {
			number0[index] += 1
			index++
		}
	}
	//println(fmt.Sprintf("%+v,%+v,%d", number0, number1, l))
	//4,14,2
	//0100
	//1110
	//0010
	for _, v := range nums {
		tmp = v
		index = 0
		for tmp > 0 {
			if tmp&1 == 1 {
				r += number0[index]
				number1[index]--
			} else {
				r += number1[index]
				number0[index]--
			}
			tmp >>= 1
			index++
		}
		for index < l {
			r += number1[index]
			number0[index]--
			index++
		}
	}

	return r*/
}
