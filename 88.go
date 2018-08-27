package main

func main() {
	println("[1,2,3,0,0,0], m = 3, [2,5,6], n = 3")
	t1 := []int{1,2,3,0,0,0, 0}
	merge(t1, 3, []int{2,5,6}, 3)
	_printArray(t1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for {
		if n == 0 {
			return
		}
		if m == 0 {
			for i := n; i > 0; i-- {
				nums1[i-1] = nums2[i-1]
			}
			return
		}
		if nums1[m-1] >= nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}

func _printArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}
