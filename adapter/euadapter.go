package main

import "fmt"

type euAdapter struct {
	humidifier *europeanHumidifier
}

func (w *euAdapter) plugInAmericanSocket() {
	fmt.Println("Adapter converts eu socket to us socket.")
	w.humidifier.plugIn()
}
