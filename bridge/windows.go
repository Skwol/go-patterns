package main

import "fmt"

type car struct {
	cable cable
}

func (c *car) charge() {
	fmt.Println("Charge request for car")
	c.cable.startCharging()
}

func (c *car) setCable(cable cable) {
	c.cable = cable
}
