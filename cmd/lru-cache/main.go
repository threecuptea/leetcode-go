package main

/**
* 146. LRU Cache
* https://leetcode.com/problems/lru-cache/description/
 */
// double linked list
type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	cache map[int]*Node
	cap   int
	left  *Node
	right *Node
}

func Constructor(capacity int) LRUCache {
	lru := &LRUCache{
		cache: make(map[int]*Node),
		cap:   capacity,
		left:  &Node{0, 0, nil, nil},
		right: &Node{0, 0, nil, nil},
	}
	lru.left.Next, lru.right.Prev = lru.right, lru.left
	return *lru
}

func (lru *LRUCache) remove(node *Node) {
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
}

// insert at th rightest
func (lru *LRUCache) insert(node *Node) {
	// Right pointer
	lru.right.Prev.Next, node.Prev = node, lru.right.Prev
	lru.right.Prev, node.Next = node, lru.right
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		// re-order
		lru.remove(node)
		lru.insert(node)
		return node.Val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key]; ok {
		// We still need to remove it; otherwise it won't re-order
		lru.remove(node)
	}
	node := &Node{key, value, nil, nil}
	lru.insert(node)
	lru.cache[key] = node
	if len(lru.cache) > lru.cap {
		delete(lru.cache, lru.left.Next.Key)
		lru.remove(lru.left.Next)
	}
}
