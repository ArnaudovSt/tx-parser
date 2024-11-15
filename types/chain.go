package types

type Chain struct {
	head *chainElement
	tail *chainElement
}

type chainElement struct {
	block *Block
	prev  *chainElement
	next  *chainElement
}

func NewChain() *Chain {
	return &Chain{}
}

func (c *Chain) Append(block *Block) {
	latest := &chainElement{block: block}

	if c.head == nil {
		c.head = latest
		c.tail = latest
	} else {
		latest.prev = c.head
		c.head.next = latest
		c.head = latest
	}
}

func (c *Chain) PeekHead() *Block {
	if c.head == nil {
		return nil
	}

	return c.head.block
}

func (c *Chain) PopHead() *Block {
	if c.head == nil {
		return nil
	}

	block := c.head.block
	c.head = c.head.prev

	return block
}

func (c *Chain) PeekTail() *Block {
	if c.tail == nil {
		return nil
	}

	return c.tail.block
}

func (c *Chain) PopTail() *Block {
	if c.tail == nil {
		return nil
	}

	block := c.tail.block
	c.tail = c.tail.next
	c.tail.prev = nil
	return block
}
