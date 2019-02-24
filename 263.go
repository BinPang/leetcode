package main

func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	for {
		if num == 1 {
			return true
		}
		if num%2 == 0 {
			num = num / 2
			continue
		}
		if num%3 == 0 {
			num = num / 3
			continue
		}
		if num%5 == 0 {
			num = num / 5
			continue
		}
		return false
	}
}
