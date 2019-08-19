package main

func main() {
	println(findComplement(458))
	println(findComplement(9))
	println(findComplement(1))
	println(findComplement(5))
}

func findComplement(num int) int {
	if num <= 0 {
		return 0 //panic
	}
	r := 0
	var s uint = 0
	for num != 0 {
		//if num&1==0 {
		//	r = 1<<s|r
		//}
		r = ((num&1)^1)<<s | r
		num >>= 1
		s += 1
	}
	return r
}
