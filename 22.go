package main

func main() {
	t0 := generateParenthesis(0)
	for _, v := range t0{
		println(v)
	}
}

func generateParenthesis(n int) []string {
	return _generateParenthesis(n, 0)
}

func _generateParenthesis(n, stack int) []string {
	if n == 0 {
		if stack > 0 {
			t := ""
			for i:=0;i<stack ;i++  {
				t = t+")"
			}
			return []string{t}
		} else {
			return []string{""}
		}
	}
	if stack < 0 {
		return nil
	}
	r := make([]string, 0)
	tmp1 := _generateParenthesis(n-1, stack+1)
	if tmp1 != nil {
		for _, v := range tmp1{
			r = append(r, "(" + v)
		}
	}
	tmp2 := _generateParenthesis(n, stack-1)
	if tmp2 != nil {
		for _, v := range tmp2{
			r = append(r, ")" + v)
		}
	}
	return r
}
