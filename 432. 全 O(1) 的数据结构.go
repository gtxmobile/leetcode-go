package main

func main() {
	
}
type AllOne struct {
	Head     *DListNode
	Tail     *DListNode
	CacheMap map[string]*DListNode
}
type DListNode struct {
	Val  int
	Key  string
	Prev *DListNode
	Next *DListNode
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	allone := AllOne{
		Head:     &DListNode{-1, -1, nil, nil},
		Tail:     &DListNode{-1, -1, nil, nil},
		CacheMap: make(map[int]*DListNode),
	}
	allone.Head.Next, allone.Tail.Prev = allone.Tail, allone.Head
	return allone
}


/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string)  {
	node, found := this.CacheMap[key]
	if found {
		node.Prev.Next, node.Next.Prev = node.Next, node.Prev
		this.InsertHead(node)
		node.Val = value
	} else {
		node := DListNode{value, key, nil, nil}
		this.InsertHead(&node)
		this.CacheMap[key] = &node
		if len(this.CacheMap) > this.Cap {
			this.DeleteTail()
			delete(this.CacheMap, this.Tail.Prev.Key)
		}
	}
}


/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string)  {

}


/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {

}


/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {

}
func (this *AllOne) InsertHead(node *DListNode) {

	node.Next, node.Prev, this.Head.Next, this.Head.Next.Prev = this.Head.Next, this.Head, node, node
}
func (this *AllOne) DeleteTail() {
	this.Tail.Prev.Prev.Next, this.Tail.Prev = this.Tail, this.Tail.Prev.Prev
}
