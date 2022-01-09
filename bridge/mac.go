package main

import "fmt"

type house struct {
	cable cable
}

func (h *house) charge() {
	fmt.Println("Charge request for house")
	h.cable.startCharging()
}

func (h *house) setCable(c cable) {
	h.cable = c
}
