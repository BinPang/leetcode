package main

func main() {
	println(minSubArrayLen(7, []int{1, 1}))
	println(minSubArrayLen(7, []int{2, 3, 1, 2, 1, 2}))
	println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func minSubArrayLen(s int, nums []int) int {
	l := len(nums)
	r := l
	total := 0
	start := 0
	isFull := false
	for k, v := range nums {
		total += v
		if total >= s {
			isFull = true
			for start < k {
				if total-nums[start] < s {
					break
				}
				total -= nums[start]
				start++
			}
			r = min(r, k-start+1)
		}
	}
	if total >= s {
		isFull = true
	}
	for start < l {
		if total-nums[start] < s {
			break
		}
		total -= nums[start]
		start++
	}
	r = min(r, l-start+1)
	if !isFull {
		return 0
	}
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
