package main

func main() {
	//@TODO may be some exception that leetcode don't know
	println("input: -. ,need:false,result:", isNumber("-."))
	return
	println("input: 4e+ ,need:false,result:", isNumber("4e+"))
	return
	println("input: . ,need:false,result:", isNumber("."))
	println("input: 1a ,need:false,result:", isNumber("1a"))
	println("input: 1 0 ,need:false,result:", isNumber("1 0"))
	println("input: 4.e10 ,need:true,result:", isNumber("4.e10"))
	println("input: 000 ,need:true,result:", isNumber("000"))
	println("input: 00. ,need:true,result:", isNumber("00."))
	println("input: 00.000 ,need:true,result:", isNumber("00.000"))
	println("input: +.75 ,need:true,result:", isNumber("+.75"))
	println("input: 12e23 ,need:true,result:", isNumber(" 12e23 "))
	println("input: e10 ,need:false,result:", isNumber("e10"))
}

func isNumber(s string) bool {
	//0-9:48~57,.:46,-:45,e:101,E:69,+:43, :32(blank)
	//trim
	l := len(s)
	if true {
		start := 0
		end := l - 1
		for {
			if start >= l || s[start] != 32 {
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
		s = s[start : end+1]
	}
	l = len(s)
	//println("after trim:", s, l)
	if l == 0 {
		return false //can omit, we confirm s is not empty
	}
	//find first e location
	eStart := -1
	for i, v := range s {
		if v == 101 || v == 69 {
			eStart = i
		}
	}
	if eStart == 0 || eStart == l-1 {
		return false
	}
	//deal before e
	if true {
		canBeginSign := true
		canPoint := true
		//println(eStart, s)
		ebs := s
		if eStart > 0 {
			ebs = s[:eStart]
		}
		//println("before e:", ebs)
		for _, v := range ebs {
			if v >= 48 && v <= 57 {

			} else if v == 46 { //.
				if !canPoint {
					return false
				} else {
					canPoint = false
				}
			} else if v == 43 || v == 45 {
				if !canBeginSign {
					return false
				}
			} else {
				return false
			}
			canBeginSign = false
		}
		//last item can't be point.@TODO maybe +.
		lebs := len(ebs)
		if lebs == 1 && (ebs[lebs-1] == 46 || ebs[lebs-1] == 43 || ebs[lebs-1] == 45) { //point
			return false
		}
		if ebs == "-." || ebs == "+." {
			return false
		}
	}
	//deal after e
	if eStart > 0 {
		eas := s[eStart+1:]
		canBeginSign := true
		//println("after e:", eas)
		for _, v := range eas {
			if v >= 48 && v <= 57 {
			} else if v == 43 || v == 45 {
				if !canBeginSign {
					return false
				}
			} else {
				return false
			}
			canBeginSign = false
		}
		leas := len(eas)
		if leas == 1 && (eas[leas-1] == 43 || eas[leas-1] == 45) { //point
			return false
		}
	}
	return true
}
