package main

import "fmt"

func main() {
	println("ok:", fmt.Sprintf("%+v", findOrder(4, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {3, 0}})))
	println("ok:", fmt.Sprintf("%+v", findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})))
	println("ok:", fmt.Sprintf("%+v", findOrder(3, [][]int{{1, 0}, {2, 0}})))
	println("ok:", fmt.Sprintf("%+v", findOrder(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))) //1->0, 2->0, 2->1
	println("no:", fmt.Sprintf("%+v", findOrder(3, [][]int{{0, 1}, {1, 2}, {2, 0}})))
	println("no:", fmt.Sprintf("%+v", findOrder(4, [][]int{{0, 1}, {1, 2}, {2, 0}})))
	println("no:", fmt.Sprintf("%+v", findOrder(4, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}})))
	println("ok:", fmt.Sprintf("%+v", findOrder(5, [][]int{{0, 1}, {1, 2}, {1, 3}, {2, 4}, {3, 4}})))
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	in := make([]int, numCourses)
	out := make([][]int, numCourses)
	for _, v := range prerequisites {
		in[v[0]] += 1
		if out[v[1]] == nil {
			out[v[1]] = []int{v[0]}
		} else {
			out[v[1]] = append(out[v[1]], v[0])
		}
	}
	//println(fmt.Sprintf("in:%+v, out:%+v", in, out))
	r := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if in[i] != 0 {
			continue
		}
		r = append(r, i)
		in[i] = -1
		if out[i] == nil { //in[i] == 0
			continue
		}
		will := out[i]
		for len(will) > 0 {
			t := will[0]
			in[t]--
			if in[t] != 0 {
				will = will[1:]
				continue
			}
			r = append(r, t)
			in[t] = -1
			if out[t] != nil {
				will = append(will, out[t]...)
			}
			will = will[1:]
		}
	}
	//println(fmt.Sprintf("%+v, %+v", in, out))
	for _, v := range in {
		if v > 0 {
			return nil
		}
	}
	return r
}
