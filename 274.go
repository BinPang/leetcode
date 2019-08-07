package main

func main() {
	println(hIndex([]int{3, 0, 6, 1, 5}))
	println(hIndex([]int{9, 9, 9, 9, 9}))
	println(hIndex([]int{2, 3, 1, 2, 2}))
}
func hIndex(citations []int) int {
	l := len(citations)
	total := 0
	arr := make([]int, l+1)
	for i := 0; i < l; i++ {
		if citations[i] >= l {
			arr[l] ++
		} else {
			arr[citations[i]]++
		}
	}
	for i := l; i >= 0; i-- {
		total += arr[i]
		if total >= i {
			return i
		}
	}
	return 0
}
