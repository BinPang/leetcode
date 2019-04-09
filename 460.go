package main

func main() {
	cache := Constructor(3)

	cache.Put(2, 2)
	cache.Put(1, 1)
	println(cache.Get(2)) // returns 2
	println(cache.Get(1)) // returns 1
	println(cache.Get(2)) // returns 2
	cache.Put(3, 3)
	cache.Put(4, 4)
	println(cache.Get(3)) // returns -1 (not found)
	println(cache.Get(2)) // returns 2.
	println(cache.Get(1)) // returns 1
	println(cache.Get(4)) // returns 4

	//cache := Constructor(2)
	//
	//cache.Put(1, 1)
	//cache.Put(2, 2)
	//println(cache.Get(1)) // returns 1
	//cache.Put(3, 3)       // evicts key 2
	//println(cache.Get(2)) // returns -1 (not found)
	//println(cache.Get(3)) // returns 3.
	//cache.Put(4, 4)       // evicts key 1.
	//println(cache.Get(1)) // returns -1 (not found)
	//println(cache.Get(3)) // returns 3
	//println(cache.Get(4)) // returns 4
}

type LFUCache struct {
	len    int
	hash   map[int]*Item
	header *Header
}

type Item struct {
	next   *Item
	prev   *Item
	header *Header
	v      int
	k      int
}
type Header struct {
	prev      *Header
	next      *Header
	dhead     *Item
	dtail     *Item
	frequency int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		len:  capacity,
		hash: make(map[int]*Item, capacity),
	}
}

func (this *LFUCache) Get(key int) int {
	if item, ok := this.hash[key]; ok {
		oh := item.header
		this.removeItem(item)
		nh := this.addNextHeader(oh)
		this.removeHeader(oh)
		this.addItem(item, nh)
		return item.v
	} else {
		return -1
	}
}

func (this *LFUCache) Put(key int, value int) {
	if this.len == 0 {
		return
	}
	var item *Item
	var ok bool
	var oh *Header
	if item, ok = this.hash[key]; ok {
		//remove item
		oh = item.header
		this.removeItem(item)
		nh := this.addNextHeader(oh)
		this.removeHeader(oh)
		this.addItem(item, nh)
	} else {
		if len(this.hash) == this.len {
			item = this.header.dtail
			oh = item.header
			this.removeItem(item)
			this.removeHeader(oh)
			delete(this.hash, item.k)
		}
		item = &Item{}
		oh = nil
		nh := this.addNextHeader(oh)
		this.addItem(item, nh)
	}
	item.k = key
	item.v = value

	this.hash[key] = item
}

//add next header
func (this *LFUCache) addNextHeader(header *Header) *Header {
	if header == nil {
		if this.header == nil {
			this.header = &Header{frequency: 1} //init
		} else {
			if this.header.frequency == 1 {
				//do nothing
			} else {
				nh := &Header{frequency: 1}
				nh.next = this.header
				this.header.prev = nh
				this.header = nh
			}
		}
		return this.header
	} else {
		if header.next == nil {
			nh := &Header{frequency: header.frequency + 1}
			header.next = nh
			nh.prev = header
		} else {
			if header.next.frequency == header.frequency+1 {
				//do nothing
			} else {
				nh := &Header{frequency: header.frequency + 1}
				nh.next = header.next
				header.next.prev = nh
				header.next = nh
				nh.prev = header
			}
		}
		return header.next
	}
}

func (this *LFUCache) removeHeader(header *Header) {
	if header == nil {
		return
	}
	if header.dhead == nil {
		if header == this.header {
			this.header = header.next
		} else {
			header.prev.next = header.next
			header.next.prev = header.prev
		}
	}
}

func (this *LFUCache) addItem(item *Item, oh *Header) {
	item.prev = nil
	item.next = nil
	if oh.dhead == nil {
		oh.dhead = item
		oh.dtail = item
	} else {
		item.next = oh.dhead
		oh.dhead.prev = item
		oh.dhead = item
	}
	item.header = oh
}

func (this *LFUCache) removeItem(item *Item) {
	if item.next == nil && item.prev == nil {
		item.header.dhead = nil
		item.header.dtail = nil
	} else if item.next == nil {
		item.prev.next = nil
		item.header.dtail = item.prev
	} else if item.prev == nil {
		item.next.prev = nil
		item.header.dhead = item.next
	} else {
		item.prev.next = item.next
		item.next.prev = item.prev
	}
	item.header = nil
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
