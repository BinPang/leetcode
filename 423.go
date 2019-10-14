package main

import (
	"strconv"
	"strings"
)

func main() {
	println(originalDigits("owoztneoer"))
	println(originalDigits("owoztneoerowoztneoer"))
	println(originalDigits("fviefuro"))
}

func originalDigits(s string) string {
	str := make([]int, 26)
	num := make([]int, 10)
	for _, v := range s {
		str[v-'a'] += 1
	}
	//zero	one	two	three four	five	six	seven	eight	nine
	data := []string{"zero", "two", "four", "six", "eight", "one", "three", "seven", "five", "nine"}
	uniq := []byte{'z', 'w', 'u', 'x', 'g', 'o', 't', 's', 'v', 'i'}
	nums := []int{0, 2, 4, 6, 8, 1, 3, 7, 5, 9}
	for i := 0; i < 10; i++ {
		num[nums[i]] = str[uniq[i]-'a']
		for _, v := range data[i] {
			str[v-'a'] -= num[nums[i]]
		}
	}
	r := ""
	for i := 0; i < 10; i++ {
		if num[i] > 0 {
			r += strings.Repeat(strconv.Itoa(i), num[i])
		}
	}
	return r
}
