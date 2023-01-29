package cache

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{Val: val}
}
