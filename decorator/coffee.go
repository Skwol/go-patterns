package main

type coffee interface {
	getPrice() int
}

type americano struct{}

func (p *americano) getPrice() int {
	return 10
}

type coffeeWithMilk struct {
	coffee coffee
}

func (c *coffeeWithMilk) getPrice() int {
	return c.coffee.getPrice() + 5
}

type coffeeWithCinamon struct {
	coffee coffee
}

func (c *coffeeWithCinamon) getPrice() int {
	return c.coffee.getPrice() + 1
}
