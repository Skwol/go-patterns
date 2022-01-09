package main

type being struct {
	name   string
	sound  string
	weight int
	height int
}

func (b *being) setName(name string) {
	b.name = name
}

func (b being) makeSound() string {
	return b.sound
}

func (b being) averageWeight() int {
	return b.weight
}

func (b being) averageHeight() int {
	return b.height
}
