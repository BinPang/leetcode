package main

func main() {
	println(canFinish(3, [][]int{{1, 0}, {1, 2}, {0, 1}}))
	println(canFinish(3, [][]int{{1, 0}, {2, 1}, {0, 2}}))
	println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	println(canFinish(2, [][]int{{1, 0}}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	tmp := make([][]int, numCourses)
	for _, v := range prerequisites {
		if len(tmp[v[0]]) == 0 {
			tmp[v[0]] = []int{v[1]}
		} else {

		}
	}

	return true
}
