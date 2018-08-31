package main

func main() {
	isNumber(" a ")
	return
	isNumber("+12.34")
	isNumber("+12.34e12")
	return
	isNumber("+12e")
	isNumber("+12")
	isNumber("0.19-eE+ abc")
}

func isNumber(s string) bool {
	//0-9:48~57,.:46,-:45,e:101,E:69,+:43, :32(blank)
	//trim
	if true {
		l := len(s)
		start := 0
		end := l-1
		for {
			if start >=l || s[start] != 32 {
				break
			}
			start++
		}
		for {
			if end < 0 || s[end] != 32 {
				break
			}
			end--
		}
		if start > end {
			return false
		}
		s = s[start:end+1]
	}
	println(s, len(s))
	//find first e location

	for i, v := range s{
		if v == 101 || v == 69 {

		}
	}
	return true
}
