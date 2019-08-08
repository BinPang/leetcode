package main

func main() {
	println(maximumGap([]int{3, 6, 9, 1}))
	println(maximumGap([]int{0, 0, 0, 100, 100, 100, 102}))
}

func maximumGap(nums []int) int {
	l := len(nums)
	if l < 2 {
		return 0
	}
	imax, imin := nums[0], nums[0]
	for _, v := range nums {
		if imax < v {
			imax = v
		}
		if imin > v {
			imin = v
		}
	}
	if imax == imin {
		return 0
	}
	maxBucket := make([]int, l+1)
	minBucket := make([]int, l+1)
	//for e := range maxBucket {
	//	maxBucket[e] = -1
	//	minBucket[e] = -1
	//}
	capacity := (imax - imin) / (l + 1)
	if (imax-imin)%(l+1) != 0 {
		capacity += 1
	}
	index := 0
	for e := range nums {
		index = (nums[e] - imin) / capacity
		if nums[e] == imax {
			index = l
		}
		if maxBucket[index] == 0 {
			maxBucket[index] = nums[e]
		} else if nums[e] > maxBucket[index] {
			maxBucket[index] = nums[e]
		}
		if minBucket[index] == 0 {
			minBucket[index] = nums[e]
		} else if nums[e] < minBucket[index] {
			minBucket[index] = nums[e]
		}
	}
	//println(fmt.Sprintf("%+v,%+v", maxBucket, minBucket))
	r := 0
	prev := maxBucket[0]
	for i := 1; i <= l; i++ {
		if maxBucket[i] != 0 {
			if r < minBucket[i]-prev {
				r = minBucket[i] - prev
			}
			prev = maxBucket[i]
		}
	}
	return r
}
