package main

func main() {
	println("input:2.00000, 10, result:", _myPow(2.00000, 10))
}


func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n > 0 {
		return _myPow(x, n)
	} else {
		return 1.0/_myPow(x, -n)
	}
}

func _myPow(x float64, n int) float64 {
	if n==0 {
		return 1.0
	}
	s := _myPow(x, n/2)
	if n%2 == 0 {
		return s*s
	} else {
		return s*s*x
	}
}
