package main

func main() {
	println("flower,flow,flight:", longestCommonPrefix([]string{"flower","flow","flight"}))
	println("dog,racecar,car:", longestCommonPrefix([]string{"dog","racecar","car"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	l := len(strs)
	r := make([]byte, 0)
	for i, v := range []byte(strs[0]){
		status := true
		for j:=1;j<l ;j++  {
			if i >= len(strs[j]) || strs[j][i] != v {
				status = false
				break
			}
		}
		if status {
			r = append(r, v)
		} else {
			break
		}
	}
	return string(r)
}
