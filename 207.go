package main

func main() {
	println(canFinish(3, [][]int{{1, 0}, {2, 1}, {0, 2}}))
	println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	println(canFinish(2, [][]int{{1, 0}}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	tmp := make([]int, numCourses)
	for e := range tmp {
		tmp[e] = -1
	}
	for _, v := range prerequisites {
		tmp[v[0]] = v[1]
		start := v[0]
		for {
			if tmp[start] == -1 {
				break
			}
			if tmp[start] == v[0] {
				return false
			}
			start = tmp[start]
		}
	}

	return true
}
