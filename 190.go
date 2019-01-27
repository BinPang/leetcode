package main

import "strconv"

func main() {
	var i uint32 = 43261596
	o := reverseBits(43261596)
	println(strconv.FormatInt(int64(i), 2) + "\n" + strconv.FormatInt(int64(o), 2))
}

func reverseBits(num uint32) uint32 {
	var r uint32 = 0
	for i := 0; i < 32; i++ {
		r = (r << 1) + (num & 1)
		num = num >> 1
	}
	return r
}
