package main

func main() {
	println(canCompleteCircuit([]int{2}, []int{2}))
	println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}

func canCompleteCircuit(gas []int, cost []int) int {
	for e := range gas {
		gas[e] = gas[e] - cost[e]
	}

	l := len(gas)
	index := 0
	total := 0
	start := 0
	for index < l {
		if gas[index] < 0 {
			index++
			continue
		}
		total = 0
		start = index
		for {
			total += gas[start]
			if total < 0 {
				break
			}
			start++
			if start == l {
				start = 0
			}
			if start == index {
				return index
			}
		}
		if start > index {
			index = start
		} else {
			index++
		}
	}

	return -1
}
