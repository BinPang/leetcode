package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1) , len(nums2)
	totalLen := len1 + len2
	indexMin, indexMax := 0, 0
	if totalLen % 2 ==0 {
		indexMin = totalLen/2-1
		indexMax = totalLen/2
	} else {
		indexMin = (totalLen-1)/2
		indexMax= indexMin
	}
	i,j,tmpValue, totalValue := 0, 0, 0, 0
	for  {
		if i <len1 && j <len2 && nums1[i] <= nums2[j] {
			tmpValue = nums1[i]
			i++
		} else if i ==len1 {
			tmpValue = nums2[j]
			j++
		} else if j == len2 {
			tmpValue = nums1[i]
			i++
		} else {
			tmpValue = nums2[j]
			j++
		}
		if i+j-1 == indexMin {
			totalValue += tmpValue
		}
		if i+j-1 == indexMax {
			totalValue += tmpValue
			break
		}
	}
	return float64(totalValue)/2
}
