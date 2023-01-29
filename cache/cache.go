package cache

func NewLRU(capacity int) LRU {
	return LRU{
		dl:       NewDoubly[any](),
		storage:  make(map[string]*Node[any], capacity),
		size:     0,
		capacity: capacity,
	}
}

func (c *LRU) Get(key string) any {
	v, ok := c.storage[key]
	if ok {
		c.dl.MoveToBack(v)
		return v.Val.(item).value
	}

	return nil
}

func (c *LRU) Put(key string, value any) {
	e, ok := c.storage[key]
	if ok {
		n := e.Val.(item)
		n.value = value
		e.Val = n
		c.dl.MoveToBack(e)
		return
	}

	if c.size >= c.capacity {
		e := c.dl.Front()
		dk := e.Val.(item).key
		c.dl.Remove(e)
		delete(c.storage, dk)
		c.size--
	}

	n := item{key: key, value: value}
	c.dl.AddAtEnd(n)
	ne := c.dl.Back()
	c.storage[key] = ne
	c.size++
}
