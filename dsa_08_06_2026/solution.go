// DSA Daily — 2026-06-08
// Problem: https://leetcode.com/problems/lru-cache/description

package dsa08062026

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func NewNode(key int, value int) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	hm := make(map[int]*Node)
	head := NewNode(0, 0)
	tail := NewNode(0, 0)
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    hm,
		head:     head,
		tail:     tail,
	}
}

func (c *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) insertAtHead(node *Node) {
	node.next = c.head.next
	node.prev = c.head
	c.head.next.prev = node
	c.head.next = node
}


func (c *LRUCache) Get(key int) int {
	if c.cache[key] == nil {
		return -1
	}

	node := c.cache[key]
	c.remove(node)
	c.insertAtHead(node)
	return node.value
}

func (c *LRUCache) Put(key int, value int) {
	if c.cache[key] != nil {
		node := c.cache[key]
		node.value = value
		c.remove(node)
		c.insertAtHead(node)
	} else {
		if len(c.cache) == c.capacity {
			delete(c.cache, c.tail.prev.key)
			c.remove(c.tail.prev)
		}
		node := NewNode(key, value)
		c.cache[key] = node
		c.insertAtHead(node)
	}
}

func solution() {
	// implement here
}

func main() {}
