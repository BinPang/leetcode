package main

import "strings"

func main() {
	println(compareVersion("7.2", "7.10"))
	println(compareVersion("7.4", "7.4.0"))
	println(compareVersion("1", "01"))
	println(compareVersion("1.0", "1.1"))
	println(compareVersion("1.0", "1.0"))
	println(compareVersion("1.0.0", "1.0.0.1"))
	println(compareVersion("1.01.01", "1.001.001"))
}

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	l1, l2 := len(v1), len(v2)
	start1, start2 := 0, 0
	for {
		if start1 == l1 || start2 == l2 {
			if start1 == l1 && start2 == l2 {
				return 0
			}
			if start1 == l1 {
				if strings.TrimLeft(v2[start2], "0") == "" {
					start2++
					continue
				} else {
					return -1
				}
			} else {
				if strings.TrimLeft(v1[start1], "0") == "" {
					start1++
					continue
				} else {
					return 1
				}
			}
		}
		if strings.TrimLeft(v1[start1], "0") == strings.TrimLeft(v2[start2], "0") {
			start1++
			start2++
			continue
		}
		if len(v1[start1]) > len(v2[start2]) {
			v2[start2] = strings.Repeat("0", len(v1[start1])-len(v2[start2])) + v2[start2]
		} else if len(v1[start1]) < len(v2[start2]) {
			v1[start1] = strings.Repeat("0", len(v2[start2])-len(v1[start1])) + v1[start1]
		}
		if v1[start1] == v2[start2] {
			start1++
			start2++
			continue
		}
		if v1[start1] > v2[start2] {
			return 1
		} else {
			return -1
		}
	}
}
