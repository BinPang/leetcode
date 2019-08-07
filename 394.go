package main

import (
	"strconv"
	"strings"
)

type stack struct {
	l int
	d []interface{}
}

func (s *stack) push(v interface{}) {
	if len(s.d) == s.l {
		s.d = append(s.d, v)
	} else {
		s.d[s.l] = v
	}
	s.l++
}
func (s *stack) pop() interface{} {
	if s.empty() {
		panic("empty stack pop panic")
	}
	s.l--
	return s.d[s.l]
}
func (s *stack) top() interface{} {
	if s.empty() {
		panic("empty stack top panic")
	}
	return s.d[s.l-1]
}
func (s *stack) empty() bool {
	return s.l == 0
}
func (s *stack) len() int {
	return s.l
}

func decodeString(s string) string {
	l := len(s)
	if s == "" {
		return ""
	}
	st := stack{d: make([]interface{}, 0)}
	start := 0
	now := 0
	number := 0
	tmpR := ""
	for {
		if now == l {
			break
		}
		if s[now] == '[' {
			now++
		} else if s[now] == ']' {
			tmpR = st.pop().(string)
			number, _ = strconv.Atoi(st.pop().(string))
			if !st.empty() && (st.top().(string)[0] < '0' || st.top().(string)[0] > '9') {
				st.push(st.pop().(string) + strings.Repeat(tmpR, number))
			} else {
				st.push(strings.Repeat(tmpR, number))
			}
			now++
		} else if s[now] >= '0' && s[now] <= '9' {
			start = now
			now++
			for {
				if s[now] == '[' {
					break
				}
				now++
			}
			//println(fmt.Sprintf("find number,%s,%+v", s[start:now], st.d[0:st.l]))
			st.push(s[start:now])
		} else {
			start = now
			now++
			for {
				if now == l {
					break
				}
				if s[now] == ']' || (s[now] >= '0' && s[now] <= '9') {
					break
				}
				now++
			}
			//find prev is string or not
			if st.empty() {
				st.push(s[start:now])
			} else {
				if st.top().(string)[0] >= '0' && st.top().(string)[0] <= '9' {
					//number, find by ]
					st.push(s[start:now])
				} else {
					st.push(st.pop().(string) + s[start:now])
				}
			}
			//println(fmt.Sprintf("find string,%s,%+v", s[start:now], st.d[0:st.l]))
		}
	}
	return st.top().(string)
}

func main() {
	println(":", "" == decodeString(""))
	println("aaa:", "aaa" == decodeString("aaa"))
	println("3[a]2[b4[F]c]:", "aaabFFFFcbFFFFc" == decodeString("3[a]2[b4[F]c]"))
	println("11[ab]:", "ababababababababababab" == decodeString("11[ab]"))
	println(decodeString("3[a2[2[c]]]"))
	println(decodeString("3[a2[c]]"))
	println(decodeString("3[a]2[bc]"))
	println(decodeString("2[abc]3[cd]ef"))
}
