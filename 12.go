package main

func main()  {
	intToRoman(23)
}

func intToRoman(num int) string {
	level := 1
	keyNumber := map[int]string{
		4:"IV",
		9:"IX",
		40:"XL",
		90:"XC",
		400:"CD",
		900:"CM",
	}
	arrR := make([]string, 0)
	for   {
		if num <= 0 {
			println(num)
			break
		}
		a := num % 10
		println(a, level)





		num = num / 10
		level+=1
	}
	return ""
}
