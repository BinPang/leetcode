package main

type Interval struct {
	Start int
	End   int
}

func main() {
	t0 := insert([]Interval{{Start: 2, End: 5}, {Start: 6, End: 7}}, Interval{Start: 0, End: 1})
	for i, v := range t0 {
		println(i, ",start", v.Start, ",end:", v.End)
	}
	println("t0 end")

	t1 := insert([]Interval{{Start: 1, End: 3}, {Start: 6, End: 9}}, Interval{Start: 2, End: 7})
	for i, v := range t1 {
		println(i, ",start", v.Start, ",end:", v.End)
	}
	println("t1 end")

	t2 := insert([]Interval{{Start: 1, End: 5}}, Interval{Start: 2, End: 3})
	for i, v := range t2 {
		println(i, ",start", v.Start, ",end:", v.End)
	}
	println("t2 end")

	t3 := insert([]Interval{{Start: 2, End: 3}}, Interval{Start: 1, End: 5})
	for i, v := range t3 {
		println(i, ",start", v.Start, ",end:", v.End)
	}
	println("t3 end")
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	r := make([]Interval, 0)
	l := len(intervals)
	if l == 0 {
		return []Interval{newInterval}
	}

	for i, v := range intervals {
		if newInterval.End < v.Start {
			r = append(r, newInterval)
			for j := i; j < l; j++ {
				r = append(r, intervals[j])
			}
			return r
		}
		if newInterval.Start > v.End {
			r = append(r, v)
			continue
		}
		if v.Start < newInterval.Start {
			newInterval.Start = v.Start
		}
		if v.End > newInterval.End {
			newInterval.End = v.End
		}
	}
	r = append(r, newInterval)
	return r
}
