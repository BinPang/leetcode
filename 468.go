package main

import "strings"

func main() {
	println(validIPAddress("192.0.0.1"))
	println(validIPAddress("20EE:FGb8:85a3:0:0:8A2E:0370:7334"))
	println(validIPAddress("2001:0db8:85a3:033:0:8A2E:0370:7334"))
}

func validIPAddress(IP string) string {
	l := len(IP)
	if l > 39 {
		return "Neither"
	}
	typ := 0
	for i := 0; i < 5; i++ {
		if i == l {
			break
		}
		if IP[i] == '.' {
			typ = 1
			break
		} else if IP[i] == ':' {
			typ = 2
			break
		}
	}
	if typ == 1 {
		v4 := strings.Split(IP, ".")
		if len(v4) != 4 {
			return "Neither"
		}
		if validIPv4(v4[0], true) && validIPv4(v4[1], false) && validIPv4(v4[2], false) && validIPv4(v4[3], false) {
			return "IPv4"
		} else {
			return "Neither"
		}
	} else if typ == 2 {
		v6 := strings.Split(IP, ":")
		if len(v6) != 8 {
			return "Neither"
		}
		if validIPv6(v6[0]) && validIPv6(v6[1]) && validIPv6(v6[2]) && validIPv6(v6[3]) &&
			validIPv6(v6[4]) && validIPv6(v6[5]) && validIPv6(v6[6]) && validIPv6(v6[7]) {
			return "IPv6"
		} else {
			return "Neither"
		}
	} else {
		return "Neither"
	}
}
func validIPv4(s string, first bool) bool {
	l := len(s)
	if l == 1 {
		if first {
			return s[0] > '0' && s[0] <= '9'
		} else {
			return s[0] >= '0' && s[0] <= '9'
		}
	} else if l == 2 {
		return s[0] > '0' && s[0] <= '9' && s[1] >= '0' && s[1] <= '9'
	} else if l == 3 {
		if s[0] == '1' {
			return s[1] >= '0' && s[1] <= '9' && s[2] >= '0' && s[2] <= '9'
		} else if s[0] == '2' {
			if s[1] >= '0' && s[1] < '5' {
				return s[2] >= '0' && s[2] <= '9'
			} else if s[1] == '5' {
				return s[2] >= '0' && s[2] <= '5'
			}
		}
	}
	return false
}
func validIPv6(s string) bool {
	l := len(s)
	if l >= 1 && l <= 4 {
		isOk := true
		for i := 0; i < l; i++ {
			if (s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'f') || (s[i] >= 'A' && s[i] <= 'F') {
			} else {
				isOk = false
				break
			}
		}
		return isOk
	}
	return false
}
