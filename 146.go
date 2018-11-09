package main

import (
	"encoding/json"
)

func main() {
	println("need:1,-1,2")
	obj := Constructor(1)
	obj.Put(2, 1)
	println(obj.Get(2))
	obj.Put(3, 2)
	println(obj.Get(2))
	println(obj.Get(3))
	return
	println("need:1,-1,-1,3,4")
	obj = Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	println(obj.Get(1)) // returns 1
	obj.Put(3, 3)       // evicts key 2
	println(obj.Get(2)) // returns -1 (not found)
	obj.Put(4, 4)       // evicts key 1
	println(obj.Get(1)) // returns -1 (not found)
	println(obj.Get(3)) // returns 3
	println(obj.Get(4)) // returns 4

	println("need:-1,-1,4")
	obj = Constructor(1)
	println(obj.Get(1))
	obj.Put(1, 4)
	println(obj.Get(2))
	println(obj.Get(1))

	return

	ts1 := `["LRUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]`
	ts2 := `[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]`
	var t1 []string
	var t2 [][]int
	json.Unmarshal([]byte(ts1), &t1)
	json.Unmarshal([]byte(ts2), &t2)
	var tl LRUCache
	for i, v := range t1 {
		if v == "LRUCache" {
			tl = Constructor(int(t2[i][0]))
			continue
		}
		if v == "put" {
			println("put:", t2[i][0], t2[i][1])
			tl.Put(t2[i][0], t2[i][1])
		}
		if v == "get" {
			println("list:", len(tl.hash))
			tmp := tl.start
			for {
				if tmp == nil {
					break
				}
				print(tmp.v, ",")
				tmp = tmp.next
			}
			println("")
			println("get:", t2[i][0])
			println(tl.Get(t2[i][0]))
		}
	}
	return

}

type LRUCache struct {
	Len   int
	hash  map[int]*Item
	start *Item //start is newest, end is oldest
	end   *Item
}

type Item struct {
	next *Item
	prev *Item
	v    int
	k    int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Len:  capacity,
		hash: make(map[int]*Item),
	}
}

func (this *LRUCache) Get(key int) int {
	//defer func() {
	//	println("get:", key, fmt.Sprintf("%+v", *this))
	//	tmp := this.start
	//	print("list:")
	//	for {
	//		if tmp == nil {
	//			break
	//		}
	//		print(tmp.v, ":", tmp, ",")
	//		tmp = tmp.next
	//	}
	//	println("")
	//}()
	if item, ok := this.hash[key]; ok {
		if this.start == item {
			//now is the newest, nothing
		} else {
			item.prev.next = item.next
			if item.next == nil {
				this.end = item.prev
				this.end.next = nil
			} else {
				item.next.prev = item.prev
			}
			item.next = this.start
			item.prev = nil
			this.start.prev = item

			this.start = item
		}
		return item.v
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	//defer func() {
	//	println("put:", key, value, fmt.Sprintf("%+v", *this))
	//	tmp := this.start
	//	print("list:")
	//	for {
	//		if tmp == nil {
	//			break
	//		}
	//		print(tmp.v, ":", tmp, ",")
	//		tmp = tmp.next
	//	}
	//	println("")
	//}()
	if item, ok := this.hash[key]; ok {
		item.v = value
		if this.start == item {
			//head, nothing to do
		} else {
			item.prev.next = item.next
			if item.next == nil {
				this.end = item.prev
				this.end.next = nil
			} else {
				item.next.prev = item.prev
			}
			item.next = this.start
			this.start.prev = item
			this.start = item
		}
	} else {
		newItem := &Item{v: value, k: key}
		this.hash[key] = newItem
		if this.start == nil {
			this.start = newItem
			this.end = newItem
		} else {
			newItem.next = this.start
			this.start.prev = newItem
			this.start = newItem
		}
		if len(this.hash) == this.Len+1 {
			delete(this.hash, this.end.k)
			this.end = this.end.prev
			this.end.next = nil
		}
	}
}

/**
 * Your LRUobj object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := println(obj.Get(key);
 * obj.Put(key,value);
 */
