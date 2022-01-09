package main

type animal interface {
	setName(string)
	makeSound() string
	averageWeight() int
	averageHeight() int
}
