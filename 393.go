package main

func main() {
	println(validUtf8([]int{197, 130, 1}))
	println(validUtf8([]int{235, 140, 4}))
}

func validUtf8(data []int) bool {
	l := len(data)
	index := 0
	for {
		if index == l {
			return true
		}
		if data[index]>>7 == 0 {
			index++
		} else if data[index]>>5 == 6 {
			if index+1 >= l {
				return false
			}
			if data[index+1]>>6 != 2 {
				return false
			}
			index += 2
		} else if data[index]>>4 == 14 {
			if index+2 >= l {
				return false
			}
			if data[index+1]>>6 != 2 || data[index+2]>>6 != 2 {
				return false
			}
			index += 3
		} else if data[index]>>3 == 30 {
			if index+3 >= l {
				return false
			}
			if data[index+1]>>6 != 2 || data[index+2]>>6 != 2 || data[index+3]>>6 != 2 {
				return false
			}
			index += 4
		} else {
			return false
		}
	}
}
