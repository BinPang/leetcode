package main

func main() {
	println(calculate("1 + 0"))
	println(calculate("1 + 1"))
	println(calculate("1 + (1-3) "))
	println(calculate(" 2-1 + 2 "))
	println(calculate("3-(4-(5-1))+2"))
	println(calculate("3-(4-(5-(4+(4-5))))+2"))
	println(calculate("(1+(4+5+2)-3)+(6+8)"))
}

func calculate(s string) int {
	l := len(s)
	r := 0
	//ignore first blank
	start := 0
	for {
		if s[start] == ' ' {
			start++
			continue
		}
		break
	}
	signArr := make([]bool, 1)
	signArr[0] = true //first item is true
	signPoint := 0
	num := 0
	preSign := true //+ is true, - is false
	for {
		//println("before:", start, string(s[start]), r, signPoint, preSign, fmt.Sprintf("%+v", signArr))
		if s[start] >= '0' && s[start] <= '9' {
			num = num*10 + int(s[start]-'0')
		} else {
			if s[start] == '(' {
				if signPoint+1 == len(signArr) {
					signArr = append(signArr, true)
				}
				if preSign == signArr[signPoint] {
					signArr[signPoint+1] = true
				} else {
					signArr[signPoint+1] = false
				}
				preSign = true
				signPoint++
				num = 0
			} else if s[start] == ' ' {
				//empty do nothing, except last one
			} else {
				if preSign == signArr[signPoint] {
					r += num
				} else {
					r -= num
				}
				if s[start] == ')' {
					signPoint--
					preSign = signArr[signPoint]
				} else if s[start] == '+' {
					preSign = true
				} else if s[start] == '-' {
					preSign = false
				} else {
					//panic
					panic("-----")
				}
				num = 0
			}
		}
		//println("after :", start, string(s[start]), r, signPoint, preSign, fmt.Sprintf("%+v", signArr))
		start++
		if start == l {
			if preSign == signArr[signPoint] {
				r += num
			} else {
				r -= num
			}
			break
		}
	}

	return r
}
