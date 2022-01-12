package main

import "fmt"

func main() {
	shoe := &sku{name: "sku"}

	observerOne := &customer{id: "1"}
	observerTwo := &customer{id: "2"}

	shoe.register(observerOne)
	shoe.register(observerTwo)

	shoe.updateAvailability()

	fmt.Println("------------------------------------------------------------------------")
	shoe.deregister(observerOne)
	shoe.updateAvailability()
}
