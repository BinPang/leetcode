package main

func main() {
	println("input:-8147483648,need:0,return:", reverse(-8147483648))
	println("input:-12300,need:-321,return:", reverse(-12300))
	println("input:-123,need:-321,return:", reverse(-123))
	println("input:1200,need:21,return:", reverse(1200))
	println("input:2147483648,need:0,return:", reverse(2147483648))
	println("input:123,need:321,return:", reverse(123))
	println("input:0,need:0,return:", reverse(0))
}

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	min := -1
	max := 1
	for i:=0; i<31;i++  {
		min *=2
	}
	for i:=0;i<30;i++{
		max *=2
	}
	max = max-1
	max = max*2+1//min=-2147483648,max=2147483647
	r := 0
	noSign := true
	if x < 0  {
		noSign = false
	}
	startZero := true
	tmp:= 0
	for {
		tmp = x % 10
		x = x /10
		if tmp == 0 && startZero {
			continue
		}

		if noSign {
			if (max - tmp)/10 < r {
				return 0
			} else {
				r = r*10+tmp
			}
		} else {
			if (min + tmp)/10 > r {
				return 0
			} else {
				r = r*10+tmp
			}
		}
		startZero = false
		if x == 0 {
			break
		}
	}
	return r
}