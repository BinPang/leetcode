package main

func main() {
	println("input:(()()((),need:4,result:", longestValidParentheses("(()()(()"))
	//return
	println("input:(()((),need:2,result:", longestValidParentheses("(()(()"))
	println("input:()((),need:2,result:", longestValidParentheses("()(()"))
	//return
	println("input:(),need:2,result:", longestValidParentheses("()"))
	println("input:)()()),need:4,result:", longestValidParentheses(")()())"))
	println("input:(()))(),need:4,result:", longestValidParentheses(")()())"))
}

func longestValidParentheses(s string) int {
	//(:40, ):41
	l := len(s)
	r := 0
	t := make([]int, l)
	stack := make([]int, l)
	point := -1
	for i, v := range s {
		//println(v, string(v))
		if v == 41 { //)
			if point >= 0 {
				t[i] = 1
				t[stack[point]] =1
				point--
			}
		} else { //(
			point++
			stack[point] = i
		}
	}
	tmpMax := 0
	for _, v := range t {
		//println(i, v)
		if v == 1 {
			tmpMax+=1
			if r < tmpMax {
				r = tmpMax
			}
		} else {
			tmpMax = 0
		}
	}
	return r
}
