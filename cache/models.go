package cache

type Doubly[T any] struct {
	Head *Node[T]
}

type Node[T any] struct {
	Val  T
	Prev *Node[T]
	Next *Node[T]
}

type item struct {
	key   string
	value any
}

type LRU struct {
	dl       *Doubly[any]
	size     int
	capacity int
	storage  map[string]*Node[any]
}
