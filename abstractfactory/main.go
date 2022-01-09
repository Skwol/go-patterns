package main

import "fmt"

func main() {
	reebokFactory, _ := getSportsFactory("reebok")
	newbalanceFactory, _ := getSportsFactory("newbalance")

	newbalanceShoe := newbalanceFactory.makeShoe()
	newbalanceShirt := newbalanceFactory.makeShirt()

	reebokShoe := reebokFactory.makeShoe()
	reebokShirt := reebokFactory.makeShirt()

	printShoeDetails(newbalanceShoe)
	printShirtDetails(newbalanceShirt)

	printShoeDetails(reebokShoe)
	printShirtDetails(reebokShirt)
}

func printShoeDetails(s sneakers) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s tshirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
