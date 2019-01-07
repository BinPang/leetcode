package main

import "fmt"

func main() {
	str := "abc"
	max := 18
	println(fmt.Sprintf(fmt.Sprintf("%%-%ds", max), str))
}
