package main

import (
	"fmt"
)

func main() {
	car := newCar()
	if err := car.startEngine(); err != nil {
		fmt.Println("error:", err.Error())
	}
	if err := car.setSpeed(50); err != nil {
		fmt.Println("error:", err.Error())
	}
	if err := car.stopEngine(); err != nil {
		fmt.Println("error:", err.Error())
	}
	if err := car.setSpeed(0); err != nil {
		fmt.Println("error:", err.Error())
	}
	if err := car.stopEngine(); err != nil {
		fmt.Println("error:", err.Error())
	}
}
