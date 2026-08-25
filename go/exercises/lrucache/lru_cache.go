package lrucache

import "sync"

type Node[K comparable, V any] struct {
	key   K
	value V
	prev  *Node[K, V]
	next  *Node[K, V]
}

type LRUCache[K comparable, V any] struct {
	capacity int
	cache    map[K]*Node[K, V]
	head     *Node[K, V]
	tail     *Node[K, V]
	mu       sync.Mutex
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	if capacity <= 0 {
		panic("Capacity must be larger than 0")
	}

	head := &Node[K, V]{}
	tail := &Node[K, V]{}

	head.next = tail
	tail.prev = head

	return &LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*Node[K, V]),
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache[K, V]) Get(key K) (V, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	node, found := l.cache[key]

	if !found {
		var zero V
		return zero, false
	}

	l.moveToFront(node)
	return node.value, true
}

func (l *LRUCache[K, V]) Put(key K, value V) {
	l.mu.Lock()
	defer l.mu.Unlock()

	node, found := l.cache[key]

	if found {
		node.value = value
		l.moveToFront(node)
		return
	}

	newNode := &Node[K, V]{
		key:   key,
		value: value,
	}

	l.cache[key] = newNode
	l.moveToFront(newNode)

	if len(l.cache) > l.capacity {
		oldest := l.removeLast()
		delete(l.cache, oldest.key)
	}
}

func (l *LRUCache[K, V]) moveToFront(node *Node[K, V]) {
	l.remove(node)
	l.addToFront(node)
}

func (l *LRUCache[K, V]) addToFront(node *Node[K, V]) {
	node.next = l.head.next
	node.prev = l.head

	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache[K, V]) remove(node *Node[K, V]) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache[K, V]) removeLast() *Node[K, V] {
	node := l.tail.prev
	l.remove(node)
	return node
}
