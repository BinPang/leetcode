package main

func main() {
	println(removeDuplicateLetters("bacb"))
	println(removeDuplicateLetters("tsbacb"))
	println(removeDuplicateLetters("bcabc"))
	println(removeDuplicateLetters("cbacdcbc")) //acdb
}

func removeDuplicateLetters(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	var i int
	var j uint8
	tmpR := make([]uint8, 0)
	lastLocation := make([]int, 26)
	for i = 0; i < 26; i++ {
		lastLocation[i] = -1
	}
	i = 0
	for i < l {
		if lastLocation[s[i]-'a'] < 0 {
			lastLocation[s[i]-'a'] = i
			tmpR = append(tmpR, s[i])
		} else {
			for j = 'a'; j < s[i]; j++ {
				if lastLocation[j-'a'] > lastLocation[s[i]-'a'] {
					tmpR[lastLocation[s[i]-'a']] = '0'
					lastLocation[s[i]-'a'] = i
					tmpR = append(tmpR, s[i])
				}
			}
		}
		i++
	}
	r := make([]uint8, 0)
	for _, v := range tmpR {
		if v != '0' {
			r = append(r, v)
		}
	}

	return string(r)
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
