package main

import "fmt"

func main() {
	println(fmt.Sprintf("%+v", majorityElement([]int{1, 2, 0, 4, 5, 0, 6, 7, 8, 9, 9, 9, 9, 9, 9, 0, 0, 0, 0, 0})))
	println(fmt.Sprintf("%+v", majorityElement([]int{2, 2, 1, 1, 1, 2, 2})))
	println(fmt.Sprintf("%+v", majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2})))
	println(fmt.Sprintf("%+v", majorityElement([]int{3, 2, 3})))
	println(fmt.Sprintf("%+v", majorityElement([]int{})))
	println(fmt.Sprintf("%+v", majorityElement([]int{1})))
}

func majorityElement(nums []int) []int {
	r := make([]int, 0)
	k1, k2, k3 := 0, 0, 0
	v1, v2, v3 := 0, 0, 0
	minus := k1

	for _, v := range nums {
		if k1 == v {
			v1++
			continue
		} else if k2 == v {
			v2++
			continue
		} else if k3 == v {
			v3++
			continue
		}
		if v1 == 0 {
			k1 = v
			v1 = 1
			continue
		} else if v2 == 0 {
			k2 = v
			v2 = 1
			continue
		} else if v3 == 0 {
			k3 = v
			v3 = 1
			continue
		}
		if minus == k1 {
			v1--
			minus = k2
		} else if minus == k2 {
			v2--
			minus = k3
		} else {
			v3--
			minus = k1
		}
	}
	//check three item is gt n/3
	v1, v2, v3 = 0, 0, 0
	for _, v := range nums {
		if k1 == v {
			v1++
		} else if k2 == v {
			v2++
		} else if k3 == v {
			v3++
		}
	}
	need := len(nums)/3 + 1
	if v1 >= need {
		r = append(r, k1)
	}
	if v2 >= need {
		r = append(r, k2)
	}
	if v3 >= need {
		r = append(r, k3)
	}
	return r
}
