package main

type cat struct {
	name string
	age  int
}

type catList struct {
	cats []*cat
}

func (l *catList) iterator(c *cat) iterator {
	return &catIterator{
		cats: l.cats,
		cat:  c,
	}
}
