package main

func main() {
	println("121,need true, return:", isPalindrome(121))
	println("1212,need true, return:", isPalindrome(1212))
	println("-121,need true, return:", isPalindrome(-121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	newX := 0
	copyX := x
	tmp := 0
	for {
		if copyX == 0 {
			break
		}
		tmp = copyX % 10
		newX = tmp+newX*10

		copyX = copyX /10
	}
	return newX == x
}

