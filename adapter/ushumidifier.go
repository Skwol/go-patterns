package main

import "fmt"

type americanHumidifier struct{}

func (m *americanHumidifier) plugInAmericanSocket() {
	fmt.Println("plug in american humidifier.")
}
