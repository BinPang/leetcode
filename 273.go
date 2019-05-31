package main

func main() {
	println(1000, numberToWords(1000))
	println(1000000, numberToWords(1000000))
	println(0, numberToWords(0))
	println(3, numberToWords(3))
	println(68, numberToWords(68))
	println(123, numberToWords(123))
	println(12345, numberToWords(12345))
	println(1234567, numberToWords(1234567))
	println(1234567999, numberToWords(1234567999))
}
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	t0 := []string{"", "Thousand", "Million", "Billion"}
	t1 := []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	t2 := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}

	r := ""
	tmpR := ""
	sep := 0

	for num != 0 {
		tmpR = ""

		tmp := num % 1000
		num = num / 1000
		if tmp == 0 {
			sep++
			continue
		}
		if tmp >= 100 {
			tmpR = t2[tmp/100-1] + " Hundred"
			tmp = tmp % 100
			if tmp != 0 {
				tmpR += " "
			}
		}
		if tmp == 0 {

		} else if tmp < 20 {
			tmpR = tmpR + t2[tmp-1]
		} else {
			tmpR = tmpR + t1[tmp/10-2]
			if tmp%10 != 0 {
				tmpR = tmpR + " " + t2[tmp%10-1]
			}
		}
		if sep == 0 {
			r = tmpR
		} else {
			if r == "" {
				r = tmpR + " " + t0[sep]
			} else {
				r = tmpR + " " + t0[sep] + " " + r
			}
		}
		sep++
	}

	return r
}
