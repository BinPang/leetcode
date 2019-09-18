package main

func main() {
	println(removeDuplicateLetters("bacb"))     //acb
	println(removeDuplicateLetters("tsbacb"))   //tsacb
	println(removeDuplicateLetters("bcabc"))    //abc
	println(removeDuplicateLetters("cbacdcbc")) //acdb
}

func removeDuplicateLetters(s string) string {
	r := make([]byte, 0)
	exist := make([]bool, 26)
	for _, v := range s {
		if exist[v-'a'] {
			for kk, vv := range r {
				if vv == byte(v) {
					if r[kk+1] < r[kk] {
						r = append(r, byte(v))
						r[kk] = '0'
					}
					break
				}
			}
		} else {
			r = append(r, byte(v))
			exist[v-'a'] = true
		}

		//println(string(v), string(r))
	}
	println(string(r))
	return ""

}

/*
给定一个仅包含小写字母的字符串，去除字符串中重复的字母，使得每个字母只出现一次。
需保证返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

示例 1:

输入: "bcabc"
输出: "abc"
示例 2:

输入: "cbacdcbc"
输出: "acdb"
*/
