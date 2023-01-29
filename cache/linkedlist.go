package cache

func (ll *Doubly[T]) Init() *Doubly[T] {
	ll.Head = &Node[T]{}
	ll.Head.Next = ll.Head
	ll.Head.Prev = ll.Head

	return ll
}

func NewDoubly[T any]() *Doubly[T] {
	return new(Doubly[T]).Init()
}

func (ll *Doubly[T]) lazyInit() {
	if ll.Head.Next == nil {
		ll.Init()
	}
}

func (ll *Doubly[T]) insert(n, at *Node[T]) *Node[T] {
	n.Prev = at
	n.Next = at.Next
	n.Prev.Next = n
	n.Next.Prev = n

	return n
}

func (ll *Doubly[T]) insertValue(val T, at *Node[T]) *Node[T] {
	return ll.insert(NewNode(val), at)
}

func (ll *Doubly[T]) AddAtEnd(val T) {
	ll.lazyInit()
	ll.insertValue(val, ll.Head.Prev)
}

func (ll *Doubly[T]) Remove(n *Node[T]) T {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	n.Next = nil
	n.Prev = nil

	return n.Val
}

func (ll *Doubly[T]) Count() int {
	var ctr int = 0

	if ll.Head.Next == nil {
		return 0
	}

	for cur := ll.Head.Next; cur != ll.Head; cur = cur.Next {
		ctr += 1
	}

	return ctr
}

func (ll *Doubly[T]) Front() *Node[T] {
	if ll.Count() == 0 {
		return nil
	}

	return ll.Head.Next
}

func (ll *Doubly[T]) Back() *Node[T] {
	if ll.Count() == 0 {
		return nil
	}

	return ll.Head.Prev
}

func (ll *Doubly[T]) MoveToBack(n *Node[T]) {
	if ll.Head.Prev == n {
		return
	}

	ll.move(n, ll.Head.Prev)
}

func (ll *Doubly[T]) move(n, at *Node[T]) {
	if n == at {
		return
	}

	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev

	n.Prev = at
	n.Next = at.Next
	n.Prev.Next = n
	n.Next.Prev = n
}
