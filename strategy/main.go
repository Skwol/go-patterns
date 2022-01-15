package main

func main() {
	navigator := newNavigator()

	navigator.setRouter(&onFootRouter{})
	navigator.calculateRoute("store", "home")
	navigator.setRouter(&byBusRouter{})
	navigator.calculateRoute("store", "home")
	navigator.setRouter(&byCarRouter{})
	navigator.calculateRoute("store", "home")
}
