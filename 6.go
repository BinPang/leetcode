package main

func main() {
	println("input:PAYPALISHIRING,need  : PINALSIGYAHRPI")
	println("input:PAYPALISHIRING,result:", convert("PAYPALISHIRING", 4))

	println("input:PAYPALISHIRING,need  : PAHNAPLSIIGYIR")
	println("input:PAYPALISHIRING,result:", convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {
	l := len(s)
	if numRows == 1 {
		return s
	}
	r := make([]byte, l)
	index := 0
	span := 0
	step := numRows + numRows - 2
	for i := 0; i < numRows; i++ {
		for j := i; j < l; j += step {
			//println("_", i, index, j, string(s[j]))
			r[index] = s[j]
			index++
			span = j + 2*(numRows-i-1)
			if i != 0 && i != numRows-1 && span < l {
				//println("_", i, index, j, string(s[span]))
				r[index] = s[span]
				index++
			}
		}
	}
	return string(r)
}
