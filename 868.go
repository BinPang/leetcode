package main

func main() {
	println(binaryGap(16410))
	println(binaryGap(26))
	println(binaryGap(22))
	println(binaryGap(5))
	println(binaryGap(6))
	println(binaryGap(8))
}

func binaryGap(N int) int {
	if N <= 0 {
		return 0
	}
	r := 0
	for {
		if N&1 == 1 {
			N >>= 1
			break
		}
		N >>= 1
		if N == 0 {
			return 0
		}
	}
	tmp := 1
	for {
		if N&1 == 1 {
			if r < tmp {
				r = tmp
			}
			tmp = 1
			N >>= 1
			continue
		}
		N >>= 1
		tmp++
		if N == 0 {
			return r
		}
	}
}
