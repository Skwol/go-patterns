package main

import "fmt"

func main() {
	coffee := &americano{}

	cappuccinoCoffee := &coffeeWithMilk{coffee: coffee}

	coffeeWithMilkAndCinamon := &coffeeWithCinamon{coffee: cappuccinoCoffee}

	fmt.Printf("Coffee with milk and cinamon price: %d\n", coffeeWithMilkAndCinamon.getPrice())
}
