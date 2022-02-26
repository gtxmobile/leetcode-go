package main

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)    // 返回  1
	cache.Put(3, 3) // 该操作会使得关键字 2 作废
	cache.Get(2)    // 返回 -1 (未找到)
	cache.Put(4, 4) // 该操作会使得关键字 1 作废
	cache.Get(1)    // 返回 -1 (未找到)
	cache.Get(3)    // 返回  3
	cache.Get(4)    // 返回  4

}

type DListNode struct {
	Val  int
	Key  int
	Prev *DListNode
	Next *DListNode
}
type LRUCache struct {
	Head     *DListNode
	Tail     *DListNode
	CacheMap map[int]*DListNode
	Cap      int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Head:     &DListNode{-1, -1, nil, nil},
		Tail:     &DListNode{-1, -1, nil, nil},
		CacheMap: make(map[int]*DListNode, capacity),
		Cap:      capacity,
	}
	lru.Head.Next, lru.Tail.Prev = lru.Tail, lru.Head
	return lru
}

func (this *LRUCache) Get(key int) int {
	node, found := this.CacheMap[key]
	if found {
		node.Prev.Next, node.Next.Prev = node.Next, node.Prev
		this.InsertHead(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
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
func (this *LRUCache) InsertHead(node *DListNode) {

	node.Next, node.Prev, this.Head.Next, this.Head.Next.Prev = this.Head.Next, this.Head, node, node
}
func (this *LRUCache) DeleteTail() {
	this.Tail.Prev.Prev.Next, this.Tail.Prev = this.Tail, this.Tail.Prev.Prev
}
