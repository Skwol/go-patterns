package main

type catIterator struct {
	index int
	cat   *cat
	cats  []*cat
}

func (c *catIterator) next() bool {
	if c.hasNext() {
		*c.cat = *c.cats[c.index]
		c.index++
		return true
	}
	return false
}

func (c *catIterator) hasNext() bool {
	if c.index < len(c.cats) {
		return true
	}
	return false
}
