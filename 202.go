package main

func isHappy(n int) bool {
	tmp := map[int]bool{}
	for {
		sum := 0
		for {
			if n < 10 {
				sum += n * n
				break
			} else {
				p := n % 10
				sum += p * p
				n = n / 10
			}
		}
		if sum == 1 {
			return true
		}
		if tmp[sum] {
			return false
		} else {
			tmp[sum] = true
		}
		n = sum
	}
}
