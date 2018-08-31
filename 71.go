package main

import "strings"

func main() {
	println("input:/a/./b///../c/../././../d/..//../e/./f/./g/././//.//h///././/..///,need:/e/f/g,result:", simplifyPath("/a/./b///../c/../././../d/..//../e/./f/./g/././//.//h///././/..///"))
	println("input:/home/of/foo/../../bar/../../is/./here/.,need:/is/here,result:", simplifyPath("/home/of/foo/../../bar/../../is/./here/."))
	println("input:/a/./b/../../c/,need:/c,result:", simplifyPath("/a/./b/../../c/"))
	println("input:/home/,need:/home,result:", simplifyPath("/home/"))
	println("input:/,need:/,result:", simplifyPath("/"))
	println("input:,need:,result:", simplifyPath(""))
	println("input:/home//foo/,need:/home/foo,result:", simplifyPath("/home//foo/"))
}

func simplifyPath(path string) string {
	ignoreBlank := make([]string, 0)
	r := ""
	start := 0
	for _, v := range strings.Split(path, "/") {
		if v == "" || v == "." {
			continue
		}
		//println(start, v)
		if v == ".." {
			start--
		} else {
			if start < 0 {
				start = 0
			}
			if start == len(ignoreBlank) {
				ignoreBlank = append(ignoreBlank, v)
				start++
			} else if start >= 0 {
				ignoreBlank[start] = v
				start++
			}
		}
	}
	for i := 0; i < start; i++ {
		r = r + "/" + ignoreBlank[i]
	}
	if r == "" {
		if path == "" {
			return ""
		} else {
			return "/"
		}
	}

	return r
}
