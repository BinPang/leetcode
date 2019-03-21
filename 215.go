package main

/*
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
 */

func main() {
	println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 2))
	println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 3))
	println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 1))
	println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 3))
	println(findKthLargest([]int{6, 2, 1, 5, 6, 4}, 3))
}

func findKthLargest(nums []int, k int) int {
	heap := make([]int, k)
	l := len(nums)
	i, j := 0, 0
	left, right := 0, 0
	for i = 0; i < k; i++ {
		j = i
		heap[j] = nums[j]
		for {
			parent := (j - 1) % 2
			if parent < 0 {
				break
			}
			if heap[parent] > heap[j] {
				heap[parent], heap[j] = heap[j], heap[parent]
			}
			j = parent
		}
	}
	for i := k; i < l; i++ {
		if nums[i] < heap[0] {
			continue
		}
		j := 0
		heap[j] = nums[i]
		for {
			left = j*2 + 1
			right = j*2 + 2
			if right > k {
				break
			} else if right == k {
				if heap[left] < heap[j] {
					heap[left], heap[j] = heap[j], heap[left]
				}
				break
			}
			if heap[left] < heap[right] {
				if heap[left] < heap[j] {
					heap[left], heap[j] = heap[j], heap[left]
				}
				j = left
			} else {
				if heap[right] < heap[j] {
					heap[right], heap[j] = heap[j], heap[right]
				}
				j = right
			}
		}
	}
	return heap[0]
}
