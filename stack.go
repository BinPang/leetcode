package main

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
