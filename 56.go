package main

import "sort"

type Interval struct {
	Start int
	End   int
}

func main() {
	t := []Interval{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	_printInterval(merge(t))

	t = []Interval{
		{1, 4},
		{4, 5},
	}
	_printInterval(merge(t))

	t = []Interval{
		{1, 4},
		{2, 3},
	}
	_printInterval(merge(t))

	_printInterval(merge(nil))

	t = []Interval{
		{4, 6},
		{3, 5},
	}
	_printInterval(merge(t))

	t = []Interval{
		{1, 4},
		{0, 4},
	}
	_printInterval(merge(t))
}

type IntervalSlice []Interval

func (a IntervalSlice) Len() int {
	return len(a)
}
func (a IntervalSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a IntervalSlice) Less(i, j int) bool {
	return a[j].Start > a[i].Start
}
func merge(intervals []Interval) []Interval {
	l := len(intervals)
	if l == 0 {
		return nil
	}
	r := make([]Interval, 0)
	sort.Sort(IntervalSlice(intervals))
	for _, v := range intervals {
		if len(r) == 0 {
			r = append(r, Interval{Start: v.Start, End: v.End})
		} else {
			if v.Start <= r[len(r)-1].End { //if not sort, use binary search
				if v.End > r[len(r)-1].End {
					r[len(r)-1].End = v.End
				}
			} else {
				r = append(r, Interval{Start: v.Start, End: v.End})
			}
		}
	}
	return r
}

func _printInterval(t []Interval) {
	for _, v := range t {
		print("[", v.Start, ",", v.End, "],")
	}
	println("")
}
