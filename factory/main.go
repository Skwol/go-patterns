package main

import "fmt"

func main() {
	dogBeing, _ := getAnimal("dog")
	catBeing, _ := getAnimal("cat")

	printDetails(dogBeing)
	printDetails(catBeing)
}

func printDetails(a animal) {
	fmt.Printf("It sounds like: %s\n", a.makeSound())
	fmt.Printf("Average weight: %d\n", a.averageWeight())
	fmt.Printf("Average height: %d\n", a.averageHeight())
}
