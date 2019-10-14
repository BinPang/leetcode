package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	t := Constructor()
	println(t.Insert(1))
	println(t.Insert(10))
	println(t.Insert(20))
	println(t.Insert(30))
	println(t.GetRandom())
	println(t.GetRandom())
	println(t.GetRandom())
	println(t.GetRandom())
	println(t.GetRandom())
	println(t.GetRandom())
}

type RandomizedSet struct {
	arr  []int
	v    map[int]int
	rand int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		arr: make([]int, 0),
		v:   make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.v[val]; ok {
		return false
	}
	this.arr = append(this.arr, val)
	this.v[val] = len(this.arr) - 1
	this.rand = 0
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.v[val]; !ok {
		return false
	}
	lastIndex := len(this.arr) - 1
	if this.v[val] == lastIndex {
		this.arr = this.arr[0:lastIndex]
	} else {
		this.v[this.arr[lastIndex]] = this.v[val]
		this.arr[this.v[val]] = this.arr[lastIndex]
		this.arr = this.arr[0:lastIndex]
	}
	delete(this.v, val)
	this.rand = 0
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
	//if this.rand == 0 {
	//	this.rand = len(this.arr)
	//}
	//index := rand.Intn(this.rand)
	//r := this.arr[index]
	//
	//if index+1 != this.rand {
	//	this.v[this.arr[index]], this.v[this.arr[this.rand-1]] = this.v[this.arr[this.rand-1]], this.v[this.arr[index]]
	//	this.arr[index], this.arr[this.rand-1] = this.arr[this.rand-1], this.arr[index]
	//}
	//
	//this.rand--
	//if this.rand == 0 {
	//	this.rand = len(this.arr)
	//}
	//
	//return r
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
