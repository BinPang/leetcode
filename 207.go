package main

func main() {
	println("true:", canFinish(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	println("true:", canFinish(3, [][]int{{1, 0}, {1, 2}, {2, 0}}))
	println("false:", canFinish(4, [][]int{{0, 1}, {3, 1}, {1, 3}, {3, 2}}))
	println("true:", canFinish(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))
	println("true:", canFinish(1, [][]int{}))
	println("true:", canFinish(2, [][]int{{1, 0}}))
	println("false:", canFinish(3, [][]int{{1, 0}, {1, 2}, {0, 1}}))
	println("false:", canFinish(3, [][]int{{1, 0}, {2, 1}, {0, 2}}))
	println("false:", canFinish(2, [][]int{{1, 0}, {0, 1}}))
	println("true:", canFinish(2, [][]int{{1, 0}}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	//prerequisites:{{1, 0}, {1, 2}, {2, 0}}
	out := make([][]int, numCourses)
	for _, v := range prerequisites {
		if out[v[0]] == nil {
			out[v[0]] = []int{v[1]}
		} else {
			out[v[0]] = append(out[v[0]], v[1])
		}
	}
	//out:{1:{0,2}, 2:{0}}
	stat := make([]int, numCourses) //0:not do, 1 doing, 2 done
	will := make([]int, 0)
	for k, v := range out { //k -> v
		if stat[k] == 2 || v == nil {
			stat[k] = 2
			continue
		}
		will = v
		stat[k] = 1
		deal := []int{k}
		for len(will) > 0 {
			t := will[0]
			if stat[t] == 1 {
				return false
			} else if stat[t] == 2 || out[t] == nil {
				stat[t] = 2
				will = will[1:]
				continue
			}
			will = will[1:]
			stat[t] = 1
			deal = append(deal, t)
			will = append(will, out[t]...)
		}
		for _, v1 := range deal {
			stat[v1] = 2
		}
	}
	return true
}
