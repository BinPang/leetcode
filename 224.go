package main

/**
Implement a basic calculator to evaluate a simple expression string.

The expression string may contain open ( and closing parentheses ), the plus + or minus sign -, non-negative integers and empty spaces .

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
type stack struct {
	data  []int
	point int
}

func (s *stack) push(data int) {
	if len(s.data) == s.point+1 {
		s.data = append(s.data, data)
		s.point++
	} else {
		s.data[s.point] = data
		s.point++
	}
}
func (s *stack) pop() int {
	if s.point > 0 {
		s.point--
		return s.data[s.point+1]
	} else {
		panic("not exist")
	}
}

func calculate(s string) int {
	l := len(s)
	r := 0
	sign := 1//1=plus, 0=minus
	st := stack{data:make([]int, 0)}
	for e := range s {
		if s[e] == ' ' {
			continue
		} else if s[e] == '(' {
			st.push(r)
			st.push(sign)

			r = 0
			sign = 1
		} else if s[e] == ')' {

		}
	}

	return r
}
