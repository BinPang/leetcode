package main

func main() {

	t2 := restoreIpAddresses("172162541")
	for _, v := range t2 {
		println(v)
	}
	return

	t1 := restoreIpAddresses("0000")
	for _, v := range t1 {
		println(v)
	}

	t0 := restoreIpAddresses("25525511135")
	for _, v := range t0 {
		println(v)
	}
}

func restoreIpAddresses(s string) []string {
	return _restoreIpAddresses(s, 4)
}

func _restoreIpAddresses(s string, num int) []string {
	l := len(s)

	if num == 1 {
		if _checkIsValid(s) {
			return []string{s}
		} else {
			return nil
		}
	}

	r := make([]string, 0)
	for i := 0; i < 3; i++ {
		if i == l {
			break
		}
		item := s[0 : i+1]
		//println("do,", s[0:i+1], s[i+1:], num)

		if check := _checkIsValid(item); !check {
			//println("**", item, len(item), check)
			continue
		}
		tmp := _restoreIpAddresses(s[i+1:], num-1)
		//println("__", item, len(tmp), s)
		if tmp != nil {
			for _, v1 := range tmp {
				r = append(r, item+"."+v1)
			}
		}
	}
	return r
}

//0~9:48~57
func _checkIsValid(s string) bool {
	//println("check,", s)
	l := len(s)
	if l == 0 || l > 3 {
		return false
	}
	if l == 1 {
		//if s[0] == 48 {
		//	return false
		//}
		return true
	}
	if s[0] == 48 {
		return false
	}
	if l == 3 {
		//255 check
		if s[0] > 50 {
			return false
		} else if s[0] == 50 {
			if s[1] > 53 || (s[1] == 53 && s[2] > 53) {
				return false
			}
		}
	}
	return true
}
