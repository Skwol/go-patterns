package main

func main() {
	lightningCable := &lightning{}
	usbcCable := &usbc{}

	houseSoc := &house{}

	houseSoc.setCable(lightningCable)
	houseSoc.charge()

	houseSoc.setCable(usbcCable)
	houseSoc.charge()

	carSoc := &car{}

	carSoc.setCable(lightningCable)
	carSoc.charge()

	carSoc.setCable(usbcCable)
	carSoc.charge()
}
