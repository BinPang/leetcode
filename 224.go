package main

import "fmt"

/**
Implement a basic calculator to evaluate a simple expression string.

The expression string may contain open ( and closing parentheses ),
the plus + or minus sign -, non-negative integers and empty spaces .

Example 1:

Input: "1 + 1"
Output: 2
Example 2:

Input: " 2-1 + 2 "
Output: 3
Example 3:

Input: "(1+(4+5+2)-3)+(6+8)"
Output: 23
Note:
You may assume that the given expression is always valid.
Do not use the eval built-in library function.
*/

func main() {
	println(calculate("3-(4-(5-1))+2"))
	return
	println(calculate("3-(4-(5-(4+(4-5))))+2"))
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
	signPoint := 1
	num := 0
	preSign := true //+ is true, - is false
	for {
		println("before:", start, string(s[start]), r, signPoint, preSign, fmt.Sprintf("%+v", signArr))
		if s[start] >= '0' && s[start] <= '9' {
			num = num*10 + int(s[start]-'0')
		} else {
			if s[start] == '(' {
				if signPoint == len(signArr) {
					signArr = append(signArr, true)
				}
				signPoint++
				if preSign {
					signArr[signPoint-1] = true
				} else {
					signArr[signPoint-1] = false
				}
				preSign = true
			} else if s[start] == ' ' {
				//empty do nothing
			} else {
				if preSign {
					r += num
				} else {
					r -= num
				}
				if s[start] == ')' {
					signPoint--
					preSign = signArr[signPoint-1]
				} else if s[start] == '+' {
					preSign = true
					if !signArr[signPoint-1] {
						preSign = false
					}
				} else if s[start] == '-' {
					preSign = false
					if !signArr[signPoint-1] {
						preSign = true
					}
				} else {
					//panic
					panic("-----")
				}
			}
			num = 0
		}
		println("after:", start, string(s[start]), r, signPoint, preSign, fmt.Sprintf("%+v", signArr))
		start++
		if start == l {
			if preSign {
				r += num
			} else {
				r -= num
			}
			break
		}
	}

	return r
}
