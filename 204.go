package main

func main() {
	println(countPrimes(10))
}

func countPrimes(n int) int {
	r := 0
	tmp := make([]bool, n)
	for i := 2; i < n; i++ {
		if tmp[i] == false {
			r++
			for j := 2; j*i < n; j++ {
				tmp[j*i] = true
			}
		}
	}
	//println(fmt.Sprintf("%+v", tmp))
	return r
}
