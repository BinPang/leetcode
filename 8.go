package main

func main() {
	println("input:42,need:42,return:", myAtoi("42"))
	println("input:-42,need:-42,return:", myAtoi("-42"))
	println("input:  -42,need:-42,return:", myAtoi("  -42"))
	println("input:- 42,need:0,return:", myAtoi("- 42"))
	println("input:2147483648,need:2147483647,return:", myAtoi("2147483648"))
	println("input:-2147483649,need:-2147483648,return:", myAtoi("-2147483649"))
	println("input:0009,need:9,return:", myAtoi("0009"))
}

func myAtoi(str string) int {
	min := -1
	max := 1
	for i := 0; i < 31; i++ {
		min *= 2
	}
	for i := 0; i < 30; i++ {
		max *= 2
	}
	max = max - 1
	max = max*2 + 1 //min=-2147483648,max=2147483647

	r := 0
	canEmpty := true
	signStatus := 0 //0:empty, 1:+, 2:-
	//0-9:48~57,-:45,+:43, :32(blank)
	for _, v := range str {
		if v == ' ' {
			if canEmpty {
				continue
			}
			break
		}
		if v == '+' || v == '-' {
			if signStatus == 0 {
				if v == '+' {
					signStatus = 1
				} else {
					signStatus = 2
				}
				canEmpty = false //if have +/-, empty is not allowed
				continue
			}
			break
		}
		if v < 48 || v > 57 {
			break
		}
		canEmpty = false
		if signStatus == 0 {
			signStatus = 1
		}

		if v == '0' && r == 0 { //first is zero
			continue
		}

		num := int(v - 48) //byte to int
		if signStatus == 1 {
			if (max-num)/10 < r {
				r = max
				break
			} else {
				r = r*10 + num
			}
		} else {
			num = -num
			if (min-num)/10 > r {
				r = min
				break
			} else {
				r = r*10 + num
			}
		}
	}
	return r
}
