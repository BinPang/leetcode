package main

import (
	"fmt"
	"strconv"
)

func main() {
	println(fmt.Sprintf("%+v", readBinaryWatch(2)))
	println(fmt.Sprintf("%+v", readBinaryWatch(1)))
}

func readBinaryWatch(num int) []string {
	r := make([]string, 0)
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if _readBinaryWatch(i)+_readBinaryWatch(j) == num {
				if j < 10 {
					r = append(r, strconv.Itoa(i)+":0"+strconv.Itoa(j))
				} else {
					r = append(r, strconv.Itoa(i)+":"+strconv.Itoa(j))
				}
			}
		}
	}
	return r
}

func _readBinaryWatch(num int) int {
	r := 0
	for num > 0 {
		if num&1 == 1 {
			r += 1
		}
		num >>= 1
	}
	return r
}
