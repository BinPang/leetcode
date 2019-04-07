package main

func main() {
	println("23:", calculate("3+4*5"))
	println("-17:", calculate("3-4*5"))
	println("103:", calculate("3+4*5*5"))
	println("103:", calculate("3 + 4  * 5  *5      "))
	println("1:", calculate(" 3/2 "))
	println("5:", calculate(" 3+5 / 2 "))
}

func calculate(s string) int {
	r := 0
	l := len(s)
	newNum := 0
	findNumber := true
	number := [2]int{}
	operater := [2]uint8{}
	pt := -1
	i := 0
	//head empty ignore
	for {
		if s[i] == ' ' {
			i++
			continue
		}
		break
	}
	for {
		if i == l {
			if pt == -1 {
				r = newNum
			} else if pt == 0 {
				r = dodo(number[0], newNum, operater[0])
			} else {
				if operater[0] == '*' || operater[0] == '/' {
					r = dodo(number[0], number[1], operater[0])
					r = dodo(r, newNum, operater[1])
				} else {
					if operater[1] == '*' || operater[1] == '/' {
						r = dodo(number[1], newNum, operater[1])
						r = dodo(number[0], r, operater[0])
					} else {
						r = dodo(number[0], number[1], operater[0])
						r = dodo(r, newNum, operater[1])
					}
				}
			}
			break
		}
		if s[i] == ' ' {
			i++
			continue
		}
		if findNumber {
			if '0' <= s[i] && s[i] <= '9' {
				newNum = newNum*10 + int(s[i]-'0')
				i++
			} else {
				findNumber = false
			}
			continue
		} else {
			//println(i, number[0], number[1], string(operater[0]), string(operater[1]), newNum)
			if pt == -1 || pt == 0 {
				pt++
				number[pt] = newNum
				operater[pt] = s[i]
			} else {
				if operater[0] == '*' || operater[0] == '/' {
					number[0] = dodo(number[0], number[1], operater[0])
					operater[0] = operater[1]
					operater[1] = s[i]
					number[1] = newNum
				} else {
					//operater[0] == '+' || '-'
					if operater[1] == '*' || operater[1] == '/' {
						number[1] = dodo(number[1], newNum, operater[1])
						operater[1] = s[i]
					} else {
						number[0] = dodo(number[0], number[1], operater[0])
						operater[0] = operater[1]
						operater[1] = s[i]
						number[1] = newNum
					}
				}
			}
			i++
			newNum = 0
			findNumber = true
		}
	}
	return r
}

func dodo(a, b int, op uint8) int {
	if op == '+' {
		return a + b
	} else if op == '-' {
		return a - b
	} else if op == '*' {
		return a * b
	} else {
		return a / b
	}
}
