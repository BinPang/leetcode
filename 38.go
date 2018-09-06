package main

import "fmt"

func main() {
	println("n=2,result:", countAndSay(2))
	println("n=3,result:", countAndSay(3))
	println("n=4,result:", countAndSay(4))
	println("n=5,result:", countAndSay(5))

}

func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}
	r := "1"
	end:= 0
	var start uint8
	count := 0
	tmpR := ""
	for {
		n--
		if n == 0 {
			break
		}
		end = len(r) -1
		start = r[end]
		count = 1
		tmpR = ""
		for {
			if end == 0 {
				tmpR = fmt.Sprint(count)+string(start)+tmpR
				//println(r, count, string(start), tmpR)
				break
			}
			if start == r[end-1] {
				count ++
			} else {
				tmpR = fmt.Sprint(count)+string(start)+tmpR
				start = r[end-1]
				count = 1
			}
			end --
		}
		r = tmpR
	}
	return r
}
