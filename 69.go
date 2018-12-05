package main

func main() {
	println(mySqrt(0))
	println(mySqrt(1))
	println(mySqrt(4))
	println(mySqrt(8))
	println(mySqrt(9))
	println(mySqrt(99)*mySqrt(99), 99, (mySqrt(99)+1)*(mySqrt(99)+1))
	println(mySqrt(5699)*mySqrt(5699), 5699, (mySqrt(5699)+1)*(mySqrt(5699)+1))
	println(mySqrt(5625)*mySqrt(5625), 5625, (mySqrt(5625)+1)*(mySqrt(5625)+1))
}

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	start := 1
	end := x
	middle := 0
	middleM := 0
	for {
		if start+1 == end {
			if end*end == x {
				return end
			} else {
				return start
			}
		}
		middle = (start + end) / 2
		middleM = middle * middle
		if middleM == x {
			return middle
		} else if middleM > x {
			end = middle
		} else {
			start = middle
		}
	}
}
