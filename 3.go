package main

func main() {
	println("zwnigfunjwz,need:8,return:", lengthOfLongestSubstring("zwnigfunjwz"))
	return
	println("wobgrovw,need:6,return:", lengthOfLongestSubstring("wobgrovw"))
	println("tmmzuxt,need:5,return:", lengthOfLongestSubstring("tmmzuxt"))
	println("abba,need:2,return:", lengthOfLongestSubstring("abba"))
	println("abbba,need:2,return:", lengthOfLongestSubstring("abbba"))
	println("aab,need:2,return:", lengthOfLongestSubstring("aab"))
	println("abb,need:2,return:", lengthOfLongestSubstring("abb"))
	println("abcabcbb,need:3,return:", lengthOfLongestSubstring("abcabcbb"))
	println("bbbbb,need:1,return:", lengthOfLongestSubstring("bbbbb"))
	println("pwwkew,need:3,return:", lengthOfLongestSubstring("pwwkew"))
	println(",need:0,return:", lengthOfLongestSubstring(""))
	println("ababab,need:2,return:", lengthOfLongestSubstring("ababab"))
}

func lengthOfLongestSubstring(s string) int {
	step := map[byte]int{}
	lastMatch := 0
	maxNow := 0
	max := 0
	for k, v := range []byte(s) {
		if lastStep, exist := step[v]; exist {
			if lastStep < lastMatch {
				maxNow = k - lastMatch + 1
				//println("--", k, lastMatch, lastStep)
			} else {
				maxNow = k - lastStep
				lastMatch = lastStep + 1
				//println("++", k, lastMatch, lastStep)
			}
		} else {
			maxNow += 1
		}
		step[v] = k
		if maxNow > max {
			max = maxNow
		}
	}
	return max
}
